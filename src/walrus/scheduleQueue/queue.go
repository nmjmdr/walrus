package scheduleQueue

import (
	"time"
)

type Queue interface {
	Add(jobType string, payload string, runAfterSeconds time.Duration) (string, error)
	Delete(jobId string) error
	Update(jobId string, payload string) error
}
