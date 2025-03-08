package dependencies

import (
	"database/sql"
	"fmt"
	"hexagonal/src/helpers"
	"hexagonal/src/products/application"
	"hexagonal/src/products/application/services"
	"hexagonal/src/products/infraestructure"
	"hexagonal/src/products/infraestructure/adapters/rabbit"
)

var (
	mySQL           infraestructure.MySQL
	db              *sql.DB
	rabbitMQAdapter *rabbit.RabbitMQAdapter
)

func Init() {
	var err error
	db, err = helpers.ConnectToDB()
	if err != nil {
		fmt.Println("Error conectando a la base de datos")
		return
	}

	mySQL = *infraestructure.NewMySQL(db)

	rabbitMQAdapter, err = rabbit.NewRabbitMQAdapter()
	if err != nil {
		fmt.Println("Error iniciando RabbitMQ:", err)
		return
	}
}

func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
	
}

func GetCreateProductController() *infraestructure.CreateProductController {
	serviceNotification := services.NewServiceNotification(rabbitMQAdapter)
	caseCreateProduct := application.NewCreateProduct(&mySQL, serviceNotification)
	return infraestructure.NewCreateProductController(caseCreateProduct)
}

func GetAllProductController() *infraestructure.GetAllProductController {
	caseGetAllProduct := application.NewGetAllProduct(&mySQL)
	return infraestructure.NewGetAllProductController(*caseGetAllProduct)
}

func GetByIdController() *infraestructure.GetByIdProductController {
	caseByIdProduct := application.NewGetByIdProduct(&mySQL)
	return infraestructure.NewGetByIdProductController(caseByIdProduct)
}

func GetDeleteController() *infraestructure.DeleteProductController {
	caseDelete := application.NewDeleteProduct(&mySQL)
	return infraestructure.NewDeleteProductController(caseDelete)
}

func GetUpdateController() *infraestructure.UpdateProductController {
	caseUpdate := application.NewUpdateProduct(&mySQL)
	return infraestructure.NewUpdateProductController(caseUpdate)
}
