package schedulerqueue

import (
	"github.com/go-redis/redis"
	"walrus/models"
  "os"
  "strings"
  "log"
  "time"
  "github.com/satori/go.uuid"
  "encoding/json"
)

const SCHEDULER_QUEUE = "SCHEDULER_QUEUE"

type redisQueue struct {
	client *redis.Client
}

func newJob(jobType string, payload string, runAt int64) models.Job {
  id := uuid.Must(uuid.NewV4())
  return models.Job { Id: id.String(), Type: jobType, Paylaod: payload, RunAt: runAt }
}

func (r *redisQueue) Add(jobType string, payload string, runAfterSeconds time.Duration) string {
  runAt := time.Now().Add(runAfterSeconds * time.Second).UnixNano()
  job := newJob(jobType, payload, runAt)

  serialized, err := json.Marshal(job)
  if err != nil {
    panic("Could not serialize job during add")
  }

  r.client.ZAdd(SCHEDULER_QUEUE, redis.Z{
    Score: float64(runAt),
    Member: string(serialized[:]),
  })
	return job.Id
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
