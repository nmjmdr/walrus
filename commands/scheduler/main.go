package main

import (
    "github.com/joho/godotenv"
    "log"
    "walrus/schedule"
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

  uErr := rq.Update(jobId, "new payload")
  fmt.Println(uErr)

  uErr = rq.Update("d9x8uad", "new payload")
  fmt.Println(uErr)

  uErr = rq.Delete("d9x8uad")
  fmt.Println(uErr)

  uErr = rq.Delete(jobId)
  fmt.Println(uErr)

}
