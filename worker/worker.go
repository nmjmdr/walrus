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
	handlers map[string]Handler
}

func copyMap(m map[string]Handler) map[string]Handler {
	if m == nil {
		return m
	}
	n := make(map[string]Handler, len(m))
	for k, v := range m {
		n[k] = v
	}
}

func NewWorker(handlers map[string]Handler) *Worker {
	w := &Worker{}
	w.handlers = copyMap(handlers)
	options := utils.LoadRedisOptions()
	d.client = redis.NewClient(options)
	return w
}
