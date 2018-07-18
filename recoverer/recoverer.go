package recoverer

// The recoverer is responsible for requeuing the jobs onto worker_queue which a particular worker has
// taken too long to process

import (
	"github.com/go-redis/redis"
	"log"
	"walrus/constants"
	"walrus/lock"
	"walrus/models"
	"walrus/utils"
)

type Recoverer struct {
	client *redis.Client
	quitCh chan bool
}

// recoverer logic:
// if a job is in PROCESSING_QUEUE but not locked then put onto worker queue

func NewRecoverer() *Recoverer {
	r := &Recoverer{}
  options := utils.LoadRedisOptions()
	r.client = redis.NewClient(options)
	r.quitCh = make(chan bool)
	return r
}

func (r *Recoverer) recover() {
	// read head of processing queue
	results, err := r.client.LRange(constants.PROCESSING_QUEUE, 0, 0).Result()
	if err != nil && err != redis.Nil {
		log.Printf("Error while trying to recover jobs in Recoverer: %s", err)
		return
	}

	if err == redis.Nil || results == nil || len(results) == 0 {
		// nothing to recover return
		return
	}

	var job *models.Job
	job, err = utils.ToJob(results[0])
	if err != nil {
    log.Printf("Error while serializing job %s, Error %s", results[0], err)
    return
	}

	// is there a active lock on the job?
	lck := lock.NewLockExp(r.client)
	isLocked, err := lck.IsLocked(job.Id)
	if err != nil {
		log.Printf("Could not check for lock in recover, Error: %s", err)
		// still attempt to recover the job
	}
	//push job to worker queue
  if isLocked {
    return
  }

	workerQueue := utils.GetWorkerQueueName(job.Type)
	_, err = r.client.RPush(workerQueue, results[0]).Result()
	if err != nil {
		log.Printf("Unable to push job %s to worker queue, Error: %s", err)
		return
	}
	_, err = r.client.LRem(constants.PROCESSING_QUEUE, 1, results[0]).Result()
	if err != nil {
		log.Printf("Unable to delete job %s from perocessing queue, Error: %s", err)
		return
	}
}

func (r *Recoverer) Start() {
	for {
		select {
		case _ = <-r.quitCh:
			break
		default:
			r.recover()
		}
	}
}

func (r *Recoverer) Stop() {
	r.quitCh <- true
}
