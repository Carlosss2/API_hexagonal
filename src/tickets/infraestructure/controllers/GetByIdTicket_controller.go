package controllers

import (
	"hexagonal/src/tickets/application"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetByIdTicketController struct {
	useCaseGetById *application.GetByIdTicket
}


func NewGetByIdTicketController(useCaseGetById *application.GetByIdTicket)*GetByIdTicketController{
	return &GetByIdTicketController{useCaseGetById: useCaseGetById}
}

func (getByIdTicket *GetByIdTicketController) GetByIdTicket(ctx *gin.Context){
	idParam := ctx.Param("id")
	id,err := strconv.Atoi(idParam)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"id invalid"})
		return
	}

	ticket,err := getByIdTicket.useCaseGetById.Execute(int32(id))

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"ticket":ticket})
}