package worker

type exampleHandler struct {
}

func (e *exampleHandler) Process(paylaod string) (string, error) {
  result := "Result: " + payload
  return result, nil
}
