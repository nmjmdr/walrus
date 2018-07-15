package main

import (
    "github.com/joho/godotenv"
    "log"
    "walrus/schedule"
    "walrus/dispatcher"
    "fmt"
    "time"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  rq := schedule.GetSchedule()

  jobId, _ := rq.Add("type1", "payload1", time.Duration(2))
  fmt.Println("Job id: ", jobId)

  d := dispatcher.NewDispatcher()
  d.Start()
}
