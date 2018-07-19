package routes

import "github.com/gorilla/mux"
import "net/http"
import "fmt"
import "walrus/schedule/api/controllers"
import "walrus/common"

func GetRoutes() []Route {
	return []RouteHandler{RouteHandler{Route: "/schedule/",
		Method:  "POST",
		Handler: controllers.Schedule().Add,
	},
		RouteHandler{Route: "/schedule/{jobId}",
			Method:  "PUT",
			Handler: controllers.Schedule().Update,
		},
		RouteHandler{Route: "/schedule/{jobId}",
			Method:  "DELETE",
			Handler: controllers.Schedule().Del,
		},
	}
}
