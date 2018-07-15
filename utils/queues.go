package utils

import (
  "fmt"
  "walrus/constants"
)

func GetWorkerQueueName(jobType string) string {
  return fmt.Sprintf("%s_%s",constants.WORKER_QUEUE_PREFIX,jobType)
}
