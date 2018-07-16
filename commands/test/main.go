package main

import (
    "github.com/joho/godotenv"
    "log"
    "walrus/schedule"
    "walrus/dispatcher"
    "fmt"
    "time"
    "walrus/worker"
    "walrus/postbox"
    "walrus/worker/exhandler"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  rq := schedule.GetSchedule()

  jobId, _ := rq.Add("type1", "payload1", time.Duration(60))
  fmt.Println("Job id: ", jobId)

  d := dispatcher.NewDispatcher()
  go d.Start()

  w := worker.NewWorker(exhandler.NewExampleHandler(), postbox.NewConsolePost())
  go w.Start()

  select {

  }
}
