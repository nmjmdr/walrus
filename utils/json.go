package utils

import (
  "encoding/json"
  "walrus/models"
)

func ToJson(object interface{}) (string, error) {
  serialized, err := json.Marshal(object)
  if err != nil {
    return "", err
  }
  js := string(serialized[:])
  return js, nil
}

func ToJob(jobJs string) (*models.Job, error) {
  job := &models.Job{}
  err := json.Unmarshal([]byte(jobJs), job)
  if err != nil {
    return nil, err
  }
  return job, nil
}
