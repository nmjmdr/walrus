package postbox

import (
  "models/job"
)

type ResultsPostbox interface {
  Post(job *models.Job, result string, err error)
}
