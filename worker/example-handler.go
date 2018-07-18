package worker

import (
  "time"
  "fmt"
)

type ExampleHandler struct {
  counter int
}

func NewExampleHandler() Handler {
  return &ExampleHandler{counter : 0}
}

func (e *ExampleHandler) Process(payload string) (string, error) {
  result := "Result: " + payload
  if e.counter == 0 {
    fmt.Println("With job sleeping: ")
   time.Sleep(20 * time.Second)
  } else {
    fmt.Println("Not sleeping now")
    e.counter++
  }
  return result, nil
}


func (e *ExampleHandler) JobType() string {
  return "type1"
}


func (e *ExampleHandler) VisiblityTimeoutTickCount() time.Duration {
  return 3 * time.Second
}
