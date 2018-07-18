package worker

import (
	"github.com/go-redis/redis"
	"walrus/utils"
	"walrus/constants"
	"walrus/postbox"
	"log"
	"fmt"
	"walrus/models"
	"walrus/lock"
	"time"
)


type Handler interface {
	Process(payload string) (string, error)
	JobType() string
	VisiblityTimeout() time.Duration
}


type Worker struct {
	client   *redis.Client
	handler Handler
	quitCh chan bool
	resultPost postbox.ResultsPostbox
}


func NewWorker(handler Handler, resultPost postbox.ResultsPostbox) *Worker {
	w := &Worker{}
	w.quitCh = make(chan bool)
	w.handler = handler
	options := utils.LoadRedisOptions()
	w.client = redis.NewClient(options)
	w.resultPost = resultPost
	return w
}

// Alternative approach to worker:
// https://redis.io/topics/distlock
/*
logic:
1. to work on a item, need to lock it:
		next = 0
		for i=0; ... (alternative try lrange key 0 -1, then try and lock one by one)
			job = lrange worker_queue 0 next
			SET job.Id my_random_value NX PX 30000
			If "SET" is successful, then lock acquired
			got a job to work, break
			else
			 next++
2. if locked on job to work then
	 process it
	 then remove lock (with random value)
	 then lrem worker_queue jobJs 1
*/
// The standard approach of implementing reliable queue using Redis is probably preferable
// because we can awlays run multiple process which move items from "processing" queue back to worker_queue


func (w *Worker) work() {
	workerQueue := utils.GetWorkerQueueName(w.handler.JobType())
	cmd := w.client.RPopLPush(workerQueue, constants.PROCESSING_QUEUE)
  result, err := cmd.Result()
	if err != nil && err != redis.Nil {
		log.Printf("Error getting jobs from worker queue: %s",err)
		return
	}

	if err == redis.Nil {
		return
	}

	var job *models.Job
	job, err = utils.ToJob(result)
	if err != nil {
		log.Print(fmt.Sprintf("Could not serialize job from queue: %s, result is: ",err, result))
		return
	}

	lck := lock.NewLockExp(w.client)
	locked, err := lck.Lock(job.Id, w.handler.VisiblityTimeout())

	if err != nil {
		log.Printf("Error encountred while trying to get for job: %s, Error: %s", job.Id, err)
		return
	}

	if !locked {
		// sone other worker has a lock on the job, return
		return
	}

	pResult, pErr := w.handler.Process(job.Payload)
	w.resultPost.Post(*job, pResult, pErr)
	jobJs, _ := utils.ToJson(job)
	lremCmd := w.client.LRem(constants.PROCESSING_QUEUE,1,jobJs)
	if lremCmd.Err() != nil {
		log.Print(fmt.Sprintf("Could not delete job id: %s from processing queue, Error: %s",job.Id, err))
	}
	lck.Unlock(job.Id)
}

func (w *Worker) Start() {
	for {
		select {
		case _ = <-w.quitCh:
			break
		default:
			w.work()
		}
	}
}

func (w *Worker) Stop() {
	w.quitCh <- true
}
