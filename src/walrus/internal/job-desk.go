package internal

import (
  "walrus/models"
)

type JobDesk interface {
  Submit(job models.Job) string
  Delete(jobId string) error
  Update(jobId string, payload string) error
}
