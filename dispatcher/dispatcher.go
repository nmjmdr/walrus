package dispatcher

import (
	"github.com/go-redis/redis"
	"walrus/constants"
	"walrus/utils"
	"log"
	"fmt"
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


func parseScheduleResult(cmdResult) (*models.Job, error) {
	if cmdResult.Err() != nil {
		return nil, cmdResult.Err()
	}
	
}

func (d *Dispatcher) fetch() {
	minScore := 0
	maxScore := time.UnixNano()
	err := client.Watch(func(tx *redis.Tx) error {
		cmdResult := tx.ZRangeByScore(constants.SCHEDULER_QUEUE, ZRangeBy{
			Min: strconv.FormatInt(minScore, 10)
			Max: strconv.FormatInt(maxScore, 10)
			Offset: 0
			Count: 1
		})
		parseScheduleResult(cmdResult)
	}, constants.SCHEDULER_QUEUE)

	if err != nil {

	}
	log.Print(fmt.Sprintf("Error: could not fetch from schedule queue",err))
}

/*
func (d *Dispatcher) fetch() {
	minScore := 0
	maxScore := time.UnixNano()
	d.client.Watch(constants.SCHEDULER_QUEUE)
	stringsSliceCmd := client.ZRangeByScore(constants.SCHEDULER_QUEUE, ZRangeBy{
		Min: strconv.FormatInt(minScore, 10)
		Max: strconv.FormatInt(maxScore, 10)
		Offset: 0
		Count: 1
	})
	if stringsSliceCmd.Err() != nil {
		d.client.Watch(constants.SCHEDULER_QUEUE)
		log.Print(fmt.Sprintf("Error: could not fetch from schedule queue",stringsSliceCmd.Err())
		return
	}
	results, err := stringsSliceCmd.Result()
	if err != nil {
		d.client.Watch(constants.SCHEDULER_QUEUE)
		log.Print(fmt.Sprintf("Error: could not fetch from schedule queue",stringsSliceCmd.Err())
		return
	}
	if len(results) == 0 {
		d.client.Watch(constants.SCHEDULER_QUEUE)
		return
	}
	job, err := utils.ToJob(results[0])
	if err != nil {
		d.client.Watch(constants.SCHEDULER_QUEUE)
		log.Print(fmt.Sprintf("Unable to unmarshal json in dispatcher, Error: ",err))
		return
	}

	workerQueue := utils.GetWorkerQueueName(job.Type)
	d.client.Multi()
}
*/


func (d *Dispatcher) dispatch() {
  for ;; {
    select {
    case _ = <- d.quitCh:
      break;
		default:
			d.fetch()
    }
  }
}

func (s *Dispatcher) Start() {
  go s.dispatch()
}

func (s *Dispatcher) Stop() {
	s.quitCh <- true
}
