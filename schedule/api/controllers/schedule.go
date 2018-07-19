package controllers

import (
  "net/http"
)

type ScheduleHandlers struct {
}

func (s *ScheduleHandlers) Add(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandlers) Del(w http.ResponseWriter, r *http.Request) {

}

func Schedule() *ScheduleHandlers {
  return &ScheduleHandlers{}
}
