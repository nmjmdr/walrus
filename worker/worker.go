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
}


func NewWorker(handler Handler) Worker {
	w := &Worker{}
	w.handler = handler
	options := utils.LoadRedisOptions()
	d.client = redis.NewClient(options)
	return w
}

func (w *Worker) Start() {

}

func (w *)
