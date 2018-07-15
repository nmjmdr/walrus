package dispatcher

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
	"walrus/constants"
	"walrus/models"
	"walrus/utils"
)

type Dispatcher struct {
	quitCh chan bool
	client *redis.Client
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{}
	options := utils.LoadRedisOptions()
	d.client = redis.NewClient(options)
	d.quitCh = make(chan bool)
	return d
}

func crush(errs []error) error {
	for i := 0; i < len(errs); i++ {
		if errs[i] != nil {
			return errs[i]
		}
	}
	return nil
}

func (d *Dispatcher) transact(tx *redis.Tx) error {
	minScore := int64(0)
	maxScore := time.Now().UnixNano()
	stringsSliceCmd := d.client.ZRangeByScore(constants.SCHEDULER_QUEUE, redis.ZRangeBy{
		Min:    strconv.FormatInt(minScore, 10),
		Max:    strconv.FormatInt(maxScore, 10),
		Offset: 0,
		Count:  1,
	})
	results, err := stringsSliceCmd.Result()
	if err != nil {
		return err
	}

	if len(results) == 0 {
		return nil
	}
	var job *models.Job
	job, err = utils.ToJob(results[0])
	if err != nil {
		return err
	}

	workerQueue := utils.GetWorkerQueueName(job.Type)
	_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
		rPushCmd := pipe.RPush(workerQueue, results[0])
		zRemCmd := pipe.ZRem(constants.SCHEDULER_QUEUE, results[0])
		hDelCmd := pipe.HDel(constants.JOBS_MAP, utils.GetJobKeyField(job.Id))
		_, err = pipe.Exec()
		return crush([]error{err, rPushCmd.Err(), zRemCmd.Err(), hDelCmd.Err()})
	})
	return err
}

func (d *Dispatcher) fetch() {
	err := d.client.Watch(d.transact)
	if err != nil {
		if err == redis.TxFailedErr {
			return
		}
		log.Print(fmt.Sprintf("Error: could not fetch from schedule queue", err))
	}
}

func (d *Dispatcher) Start() {
	for {
		select {
		case _ = <-d.quitCh:
			break
		default:
			d.fetch()
		}
	}
}

func (s *Dispatcher) Stop() {
	s.quitCh <- true
}
