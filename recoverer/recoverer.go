package recoverer

// The recoverer is responsible for requeuing the jobs onto worker_queue which a particular worker has
// taken too long to process

import (
	"github.com/go-redis/redis"
	"walrus/utils"
	"walrus/constants"
	"log"
	"fmt"
	"walrus/models"
)

type Recoverer struct {
  client   *redis.Client
  quitCh chan bool
}

func NewRecoverer() *Recoverer {
  r := &Recoverer{}
  return r
}

func (r *Recoverer) Start() {

}

func (r *Recoverer) Stop() {

}
