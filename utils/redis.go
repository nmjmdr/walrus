package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
	"strings"
	"walrus/constants"
)

func LoadRedisOptions() *redis.Options {
	addr := os.Getenv("REDIS_ADDR")
	if len(strings.TrimSpace(addr)) == 0 {
		log.Print("Warning: REDIS_ADDR is not defined in environment variables; connecting to default")
		addr = "localhost:6379"
	}

	password := os.Getenv("REDIS_PASSWORD")

	return &redis.Options{Addr: addr, Password: password, DB: 0}
}

func GetWorkerQueueName(jobType string) string {
	return fmt.Sprintf("%s_%s", constants.WORKER_QUEUE_PREFIX, jobType)
}

func GetJobKeyField(jobId string) string {
	return fmt.Sprintf("%s_%s", constants.JOB_FIELD_KEY_PREFIX, jobId)
}
