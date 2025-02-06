package main

import (
	product "hexagonal/src/products/infraestructure/dependencies"
	routesProduct "hexagonal/src/products/infraestructure/routes"
	ticket "hexagonal/src/tickets/infraestructure/dependencies"
	routesTickets "hexagonal/src/tickets/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	product.Init()
	ticket.Init()

	defer product.CloseDB()
	defer ticket.CloseDB()

	r := gin.Default()

	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	
	routesProduct.Routes(r)
	routesTickets.Routes(r)

	
	r.Run()
}
