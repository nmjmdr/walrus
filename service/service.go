package service

import (
	"walrus/common"
	"github.com/gorilla/mux"
  "fmt"
  "net/http"
)

type Service struct {
	muxRouter *mux.Router
	routes    []common.Route
	quitCh    chan bool
}

func NewService(routes []common.Route) *Service {
	s := &Service{}
	s.muxRouter = mux.NewRouter()
	s.routes = routes
	addRoutes(s.muxRouter, s.routes)
	s.quitCh = make(chan bool)
	return s
}

func addRoutes(muxRouter *mux.Router, routes []common.Route) {
	for _, route := range routes {
		muxRouter.HandleFunc(route.Route, route.Handler).Methods(route.Method)
	}
}

func (s *Service) Start(service string, listenAddress string) {
	go func() {
		if err := http.ListenAndServe(listenAddress, s.muxRouter); err != nil {
			panic(fmt.Sprintf("Service %s failed to start, Error: ", service, err))
		}
	}()
	s.muxRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
    tpl, err1 := route.GetPathTemplate()
    met, err2 := route.GetMethods()
    fmt.Println(tpl, err1, met, err2)
    return nil
	})
	fmt.Printf("%s listening on %s, ready to serve requests\n", service, listenAddress)

	//TO DO: not done yet: graceful handling of service being shutdown when service requests are in pipeline
	select {
	case _ = <-s.quitCh:
		fmt.Printf("%s received stop signal, stopping\n")
	}
}

func (s *Service) Stop() {
	s.quitCh <- true
}
