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
    "walrus/recoverer"
)

func main() {

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  rq := schedule.GetSchedule()

  jobId, _ := rq.Add("type1", "payload1", 2 * time.Second)
  fmt.Println("Job id: ", jobId)

  d := dispatcher.NewDispatcher()
  go d.Start()

  w := worker.NewWorker(worker.NewExampleHandler(), postbox.NewConsolePost())
  go w.Start()


  r := recoverer.NewRecoverer()
  go r.Start()


  select {

  }
}
