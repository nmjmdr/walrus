package postbox

import (
  "walrus/models"
)

type ResultsPostbox interface {
  Post(job models.Job, result string, err error)
}
