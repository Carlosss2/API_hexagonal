package routes

import (
	
	"hexagonal/src/products/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/products")
	createProduct := dependencies.GetCreateProductController().Create
	viewProduct := dependencies.GetAllProductController().View
	viewProductById := dependencies.GetByIdController().GetByIdProduct
	deleteProduct := dependencies.GetDeleteController().DeleteProduct
	updateProduct := dependencies.GetUpdateController().Update

	routes.POST("", createProduct)
	routes.GET("/getAllProducts",viewProduct)
	routes.GET("/getByIdProduct/:id",viewProductById)
	routes.DELETE("/deleteProduct/:id",deleteProduct)
	routes.PUT("/updateProduct/:id",updateProduct)
}




