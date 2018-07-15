package worker

import (
	"github.com/go-redis/redis"
	"walrus/utils"
)

type Handler interface {
	process(payload string) (string, error)
}


type Worker struct {
	client   *redis.Client
	handler Handler
	quitCh chan bool
}


func NewWorker(handler Handler) Worker {
	w := &Worker{}
	w.quitCh = make(chan bool)
	w.handler = handler
	options := utils.LoadRedisOptions()
	d.client = redis.NewClient(options)
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

}

func (w *Worker) Start() {
	for {
		select {
		case _ = <-d.quitCh:
			break
		default:
			d.work()
		}
	}
}

func (w *Worker) Stop() {

}
