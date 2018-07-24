package requests

import (
	"net/http"
	"encoding/json"
	"strings"
	"errors"
	"io/ioutil"
  )

type AddJobRequest struct {
	Payload         string `json:"payload"`
	JobType         string `json:"job_type"`
	RunAfterSecs    int    `"json:"run_after_secs"`
}


func readRequest(r *http.Request) (*AddJobRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	  defer r.Body.Close()
	  if err != nil {
		  return nil, err
	  }
	var addJobReq AddJobRequest
	err = json.Unmarshal(b, &addJobReq)
	return &addJobReq, err
  
  }
  
  func ToAddJobRequest(r *http.Request) (*AddJobRequest, error) {
	addJobReq, err := readRequest(r)
	if err != nil {
	  return nil, err
	}
	addJobReq.JobType = strings.TrimSpace(addJobReq.JobType)
	if len(addJobReq.JobType) == 0 {
	  return nil, errors.New("JobType cannot be empty")
	}
	return addJobReq, nil
}
