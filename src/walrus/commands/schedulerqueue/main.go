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

 fmt.Println(schedulerqueue.NewRedisQueue().Add("type1", "payload1", time.Duration(2)))
}
