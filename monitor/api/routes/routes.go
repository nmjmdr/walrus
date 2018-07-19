package routes

import "github.com/gorilla/mux"
import "net/http"
import "fmt"
import "walrus/monitor/api/controllers"
import "walrus/common"

func GetRoutes() []Route {
	return []RouteHandler{RouteHandler{Route: "/monitor/queues/status",
		Method:  "GET",
		Handler: controllers.Queues().Status,
	}}
}
