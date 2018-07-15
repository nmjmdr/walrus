package utils

import (
  "fmt"
  "walrus/constants"
)

func GetWorkerQueueName(jobType string) string {
  return fmt.Sprintf("%s_%s",constants.WORKER_QUEUE_PREFIX,jobType)
}


func GetJobKeyField(jobId string) string {
  return fmt.Sprintf("%s_%s",constants.JOB_FIELD_KEY_PREFIX,jobId)
}
