package main

import (
	"flag"
	"github.com/fthyilmz/workshop-go.git/app"
	"github.com/fthyilmz/workshop-go.git/config"
)

// @title Home Inventory Track API
// @version 1.0

// @contact.name API Support
// @contact.url http://fatih.im
// @contact.email hi@fatih.im

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /v1
func main() {
	addr := flag.String("addr", config.Server.Addr, "Address to listen and serve")

	flag.Parse()

	r := router.New()

	r.Run(*addr)
}
