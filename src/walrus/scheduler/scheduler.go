import (
	"github.com/go-redis/redis"
)


type Scheduler struct {
  client *redis.Client
}

func NewScheduler() *Scheduler {
  s := &Scheduler()
  options := utils.LoadRedisOptions()
	s.client = redis.NewClient(options)
  return s
}

func (s *Scheduler) Start() {

}

func (s *Scheduler) Stop() {

}
