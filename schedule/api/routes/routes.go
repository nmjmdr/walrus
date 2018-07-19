package routes

import "walrus/schedule/api/controllers"
import "walrus/common"

func GetRoutes() []common.Route {
	return []common.Route{common.Route{Route: "/schedule/",
		Method:  "POST",
		Handler: controllers.Schedule().Add,
	},
		common.Route{Route: "/schedule/{jobId}",
			Method:  "PUT",
			Handler: controllers.Schedule().Update,
		},
		common.Route{Route: "/schedule/{jobId}",
			Method:  "DELETE",
			Handler: controllers.Schedule().Del,
		},
	}
}
