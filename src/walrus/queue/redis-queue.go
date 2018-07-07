package queue

import (
	"github.com/go-redis/redis"
	"walrus/models"
  "os"
  "strings"
  "log"
)

type redisQueue struct {
	client *redis.Client
}

func (r *redisQueue) Add(job models.Job) string {
	return ""
}

func (r *redisQueue) Update(jobId string, paylaod string) error {
	return nil
}

func (r *redisQueue) Delete(jobId string) error {
	return nil
}

type redisOptions struct {
  Addr string
  Password string
  DB int
}

func loadRedisOptions() *redis.Options {
  addr := os.Getenv("REDIS_ADDR")
  if len(strings.TrimSpace(addr)) == 0 {
    log.Print("Warning: REDIS_ADDR is not defined in environment variables connecting to default")
    addr = "localhost:6379"
  }

  password := os.Getenv("REDIS_PASSWORD")

  return &redis.Options{ Addr: addr, Password: password, DB: 0 }
}

func NewRedisQueue() Queue {

	r := redisQueue{}
  options := loadRedisOptions()
	r.client = redis.NewClient(options)

	return &r
}
