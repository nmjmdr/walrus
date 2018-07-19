package controllers

type QueueHandlers struct {
}

func (s *QueueHandlers) Status(w http.ResponseWriter, r *http.Request) {

}

func Queues() *QueueHandlers {
  return &QueueHandlers{}
}
