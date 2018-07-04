package internal

import (
  "walrus/models"
)

type redisJobDesk struct {
}

func NewJobDesk() JobDesk {
  r := redisJobDesk{}
  return &r
}

func (r *redisJobDesk) Submit(job models.Job) string {
  return ""
}

func (r *redisJobDesk) Delete(jobId string) error {
  return nil
}

func (r *redisJobDesk) Update(jobId string, payload string) error {
  return nil
}
