package main

import (
    "github.com/joho/godotenv"
    "log"
    "walrus/queue"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

 queue.NewRedisQueue()
}
