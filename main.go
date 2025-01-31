package main

import (
	product "hexagonal/src/products/infraestructure/dependencies"
	routesProduct "hexagonal/src/products/infraestructure/routes"
	ticket "hexagonal/src/tickets/infraestructure/dependencies"
	routesTickets "hexagonal/src/tickets/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	product.Init()
	ticket.Init()

	defer product.CloseDB()
	defer ticket.CloseDB()

	r := gin.Default()
	routesProduct.Routes(r)
	routesTickets.Routes(r)
	r.Run()
	
	
}