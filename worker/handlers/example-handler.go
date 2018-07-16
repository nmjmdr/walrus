package exhandler

import (
  "walrus/worker"
)

type ExampleHandler struct {
}

func NewExampleHandler() worker.Handler {
  return &ExampleHandler{}
}

func (e *ExampleHandler) Process(paylaod string) (string, error) {
  result := "Result: " + payload
  return result, nil
}


func (e *ExampleHandler) JobType() string {
  return "type1"
}
