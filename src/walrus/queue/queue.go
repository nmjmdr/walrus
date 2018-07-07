package queue

import (
	"walrus/models"
)

type Queue interface {
	Add(job models.Job) string
	Delete(jobId string) error
	Update(jobId string, payload string) error
}
