package requests

import (
	"net/http"
	"encoding/json"
	"strings"
	"errors"
  )

type AddJobRequest struct {
	Payload         string `json:"payload"`
	JobType         string `json:"job_type"`
	RunAfterSecs    int    `"json:"run_after_secs"`
}

  
func ToAddJobRequest(r *http.Request) (*AddJobRequest, error) {
	var addJobReq AddJobRequest
	readBytes, err := readRequest(r)
	err = json.Unmarshal(readBytes, &addJobReq)
	if err != nil {
	  return nil, err
	}
	addJobReq.JobType = strings.TrimSpace(addJobReq.JobType)
	if len(addJobReq.JobType) == 0 {
	  return nil, errors.New("JobType cannot be empty")
	}
	return &addJobReq, nil
}
