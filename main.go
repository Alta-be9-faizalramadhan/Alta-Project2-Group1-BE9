package main

import (
	"altaproject/config"
	"altaproject/factory"
	"altaproject/middlewares"
	"altaproject/routes"
)

func main() {
	//initiate db connection
	dbConn := config.InitDB()

	//initiate factory
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":5000"))
}
