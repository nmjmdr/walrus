package main

import (
	scheduleRoutes "walrus/schedule/api/routes"
	"walrus/service"
)

func main() {
	routes := scheduleRoutes.GetRoutes()
	service := service.NewService(routes)
	service.Start("schedule-api", ":8090")
	select {}
}
