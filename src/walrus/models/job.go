package models

type Job struct {
  Id string `json:"id"`
	Type    string `json:"type"`
	Paylaod string `json:"payload"`
  RunAt int64 `json:"runAt"`
}
