import (
	"github.com/go-redis/redis"
)


type Dispatcher struct {
  client *redis.Client
  quitCh chan bool
}

func NewDispatcher() *Dispatcher {
  s := &Scheduler()
  options := utils.LoadRedisOptions()
	s.client = redis.NewClient(options)
  s.quitCh = make(chan bool)
  return s
}



func (s *Dispatcher) dispatch() {
  for ;; {
    select {
    case _ = <- s.quitCh:
      break;
    default:
      // pick jobs to dispatch here
    }
  }
}

func (s *Dispatcher) Start() {
  go s.dispatch()

}

func (s *Dispatcher) Stop() {

}
