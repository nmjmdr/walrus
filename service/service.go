package service

import (
	"common"
	"github.com/gorilla/mux"
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
	return r
}

func addRoutes(muxRouter *mux.Router, routes []common.Route) {
	for _, route := range routes {
		muxRouter.HandleFunc(route.Route, route.Handler).Method(route.Method)
	}
}

func (s *Service) Start(service string, listenAddress string) {
	go func() {
		if err := http.ListenAndServe(listenAddress, s.muxRouter); err != nil {
			log.Fatalf("Service %s failed to start, Error: ", service, err)
		} else {
			fmt.Printf("'%s listening on %s...\n", service, listenAddress)
		}
	}()

	//TO DO: not done yet: graceful handling of services being served when service attempts to quit
	select {
	case _ = <-s.quitCh:
		fmt.Printf("%s received stop signal, stopping\n")
	}
}

func (s *Service) Stop() {
	s.quitCh <- true
}
