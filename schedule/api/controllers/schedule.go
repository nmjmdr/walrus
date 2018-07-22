package controllers

import (
  "net/http"
  "walrus/schedule/api/requestparsers/addjobparser"
  "walrus/schedule"
  "time"
  "encoding/json"
)

type ScheduleHandlers struct {
  schedule schedule.Schedule
}

type AddJobResponse struct {
  JobId string `json:"job_id"`
}


func (s *ScheduleHandlers) Add(w http.ResponseWriter, r *http.Request) {
  addJobReq, err := addjobparser.ParseRequest(r)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  jobId, err := s.schedule.Add(addJobReq.JobType, addJobReq.Payload,time.Duration(addJobReq.RunAfterSecs) * time.Second)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  addJobRes := AddJobResponse{ JobId: jobId }
  output, _ := json.Marshal(addJobRes)
  w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func (s *ScheduleHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandlers) Del(w http.ResponseWriter, r *http.Request) {

}

func Schedule() *ScheduleHandlers {
  s := &ScheduleHandlers{}
  s.schedule = schedule.GetSchedule()
  return s
}
