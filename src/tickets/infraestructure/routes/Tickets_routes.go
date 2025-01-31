package routes

import (
	
	"hexagonal/src/tickets/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/tickets")
	createTicket := dependencies.GetCreateTicketController().Create
	getAllTickets := dependencies.GetGetAllTicketController().View
	deleteTickets := dependencies.GetDeleteTicketController().Delete
	updateTickets := dependencies.GetUpdateTicketController().Update


	routes.POST("/postTicket",createTicket)
	routes.GET("/getAllTickets",getAllTickets)
	routes.DELETE("/deleteTicket/:id",deleteTickets)
	routes.PUT("/updateTicket/:id",updateTickets)
}