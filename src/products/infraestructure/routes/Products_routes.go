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

	routes.POST("", createProduct)        // POST /products
	routes.GET("/", viewProduct)          // GET /products
	routes.GET("/:id", viewProductById)   // GET /products/:id
	routes.DELETE("/:id", deleteProduct)  // DELETE /products/:id
	routes.PUT("/:id", updateProduct)     // PUT /products/:id

}




