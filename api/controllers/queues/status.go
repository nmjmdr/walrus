package queues

import (
  "net/http"
  "fmt"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
  // handle queue status here later
  w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
