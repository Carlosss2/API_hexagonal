package dependencies

import (
	"database/sql"
	"fmt"
	"hexagonal/src/helpers"
	"hexagonal/src/products/application"
	"hexagonal/src/products/infraestructure"
)
var(
	mySQL infraestructure.MySQL
	db    *sql.DB
)

func Init(){
	db, err := helpers.ConnectToDB()

	if err != nil{
		fmt.Println("Server error")
		return
	}

	
	

	mySQL = *infraestructure.NewMySQL(db)
	

	
}


func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}


func GetCreateProductController()*infraestructure.CreateProductController{
	caseCreateProduct := application.NewCreateProduct(&mySQL)
	return infraestructure.NewCreateProductController(caseCreateProduct)
}

func GetAllProductController()*infraestructure.GetAllProductController{
	caseGetAllProduct := application.NewGetAllProduct(&mySQL)
	return infraestructure.NewGetAllProductController(*caseGetAllProduct)
}
func GetByIdController()*infraestructure.GetByIdProductController{
	caseByIdProduct := application.NewGetByIdProduct(&mySQL)
	return infraestructure.NewGetByIdProductController(caseByIdProduct)
}

func GetDeleteController()*infraestructure.DeleteProductController{
	caseDelete := application.NewDeleteProduct(&mySQL)
	return infraestructure.NewDeleteProductController(caseDelete)
}

func GetUpdateController()*infraestructure.UpdateProductController{
	caseUpdate := application.NewUpdateProduct(&mySQL)
	return infraestructure.NewUpdateProductController(caseUpdate)
}