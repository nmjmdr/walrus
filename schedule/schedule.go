package schedule

import (
	"time"
)

type Schedule interface {
	Add(jobType string, payload string, runAfter time.Duration) (string, error)
	Delete(jobId string) error
	Update(jobId string, payload string) error
}

var s = newRedisScheduleQueue()

func GetSchedule() Schedule {
	return s
}
