package dispacther

import (
	"walrus/schedule"
)


type Dispatcher struct {
  quitCh chan bool
}

func NewDispatcher() *Dispatcher {
  d := &Dispatcher{}
  d.quitCh = make(chan bool)
  return s
}



func (s *Dispatcher) dispatch() {
  for ;; {
    select {
    case _ = <- s.quitCh:
      break;
    default:
      // pick jobs to dispatch here from schedule queue
    }
  }
}

func (s *Dispatcher) Start() {
  go s.dispatch()

}

func (s *Dispatcher) Stop() {

}
