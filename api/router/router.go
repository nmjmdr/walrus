package router

import "github.com/gorilla/mux"
import "net/http"
import "fmt"
import "walrus/api/controllers/schedule"
import "walrus/api/controllers/queues"

func Start(listenAddress string) {
  r := mux.NewRouter()
  r.HandleFunc("/queues/status", queues.StatusHandler).Methods("GET")
  r.HandleFunc("/schedule",schedule.AddHandler).Methods("POST")
  r.HandleFunc("/schedule/{jobId}",schedule.UpdateHandler).Methods("PUT")
  r.HandleFunc("/schedule/{jobId}",schedule.DeleteHandler).Methods("DELETE")

  go func() {
		if err := http.ListenAndServe(listenAddress, r); err != nil {
			fmt.Println(err)
		} else {
      fmt.Println("Listening on port: 8090")
    }
	}()
}
