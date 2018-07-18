package worker

import (
  "time"
  "fmt"
)

type ExampleHandler struct {
  counter int
}

var e = ExampleHandler{}

func NewExampleHandler() Handler {
  return &e
}

func (e *ExampleHandler) Process(payload string) (string, error) {
  result := "Result: " + payload
  if e.counter == 0 {
    e.counter++
    fmt.Println("With job sleeping")
    time.Sleep(20 * time.Second)
    return "Just slept", nil
  } else {
    fmt.Println("Not sleeping now")
    return result, nil
  }

}


func (e *ExampleHandler) JobType() string {
  return "type1"
}


func (e *ExampleHandler) VisiblityTimeoutTickCount() time.Duration {
  return 3 * time.Second
}
