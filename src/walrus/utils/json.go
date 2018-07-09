package utils

import (
  "encoding/json"
)

func ToJson(object interface{}) (string, error) {
  serialized, err := json.Marshal(object)
  if err != nil {
    return "", err
  }
  js := string(serialized[:])
  return js, nil
}
