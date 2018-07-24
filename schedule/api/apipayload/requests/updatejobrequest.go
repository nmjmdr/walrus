package requests

import (
	"net/http"
	"encoding/json"
	"strings"
	"errors"
  )

type UpdateJobRequest struct {
	Payload         string `json:"payload"`
}

  
func ToUpdateJobRequest(r *http.Request) (*UpdateJobRequest, error) {
	var updateJobReq UpdateJobRequest
	readBytes, err := readRequest(r)
	err = json.Unmarshal(readBytes, &updateJobReq)
	if err != nil {
	  return nil, err
	}
	return &updateJobReq, nil
}
