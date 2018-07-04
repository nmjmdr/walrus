package engine

type Desk interface {
  Submit(job Job) string
  Delete(jobId string)
  Update(jobId string, payload string)
}
