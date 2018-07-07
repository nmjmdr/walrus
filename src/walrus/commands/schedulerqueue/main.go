package main

import (
    "github.com/joho/godotenv"
    "log"
    "walrus/schedulerqueue"
    "fmt"
    "time"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

 fmt.Println(schedulerqueue.NewRedisQueue().Add("type", "payload", time.Duration(2)))
}
