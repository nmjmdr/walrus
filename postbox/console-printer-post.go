package postbox

import (
  "walrus/models"
  "fmt"
)

type ConsolePost struct {
}

func New() ResultsPostbox {
  return &Console{}
}

func (c *ConsolePost) Post(job models.Job, result string, err error) {
  if err != nil {
    fmt.Printf("Job: %s posted an error response: %s",err)
    return
  }
  fmt.Printf("Job: %s posted result: %s",result)
}
