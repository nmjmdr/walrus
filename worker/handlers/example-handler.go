package worker

type exampleHandler struct {
}

func (e *exampleHandler) process(paylaod string) (string, error) {
  result := "Result: " + payload
  return result, nil
}
