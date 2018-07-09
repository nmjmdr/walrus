package utils

import (
	"github.com/go-redis/redis"
  "os"
  "strings"
  "log"
)

func LoadRedisOptions() *redis.Options {
  addr := os.Getenv("REDIS_ADDR")
  if len(strings.TrimSpace(addr)) == 0 {
    log.Print("Warning: REDIS_ADDR is not defined in environment variables connecting to default")
    addr = "localhost:6379"
  }

  password := os.Getenv("REDIS_PASSWORD")

  return &redis.Options{ Addr: addr, Password: password, DB: 0 }
}
