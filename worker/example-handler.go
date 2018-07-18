package worker

import (
  "time"
)

type ExampleHandler struct {
}

func NewExampleHandler() Handler {
  return &ExampleHandler{}
}

func (e *ExampleHandler) Process(payload string) (string, error) {
  result := "Result: " + payload
  return result, nil
}


func (e *ExampleHandler) JobType() string {
  return "type1"
}


func (e *ExampleHandler) VisiblityTimeoutTickCount() time.Duration {
  return 10 * time.Second
}
