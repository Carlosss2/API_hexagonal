package infraestructure

import (
	"hexagonal/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByIdProductController struct {
	useCaseGetById *application.GetByIdProduct
}

func NewGetByIdProductController(useCaseGetById *application.GetByIdProduct)*GetByIdProductController{
	return &GetByIdProductController{useCaseGetById: useCaseGetById}
}

func (getByIdProduct *GetByIdProductController) GetByIdProduct(ctx *gin.Context){
	idParam := ctx.Param("id")
	id,err := strconv.Atoi(idParam)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"id invalid"})
		return
	}
	product,err := getByIdProduct.useCaseGetById.Execute(int32(id))
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"product":product})
}