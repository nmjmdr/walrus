package internal

import (
  "walrus/models"
  "github.com/go-redis/redis"
)

type redisJobDesk struct {
  client redis.Client
}

func NewJobDesk() JobDesk {
  r := redisJobDesk{}
  // read from env
  //r.client =
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
