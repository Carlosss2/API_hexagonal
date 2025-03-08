package application

import (
	"hexagonal/src/products/application/services"
	"hexagonal/src/products/domain"
)

type CreateProduct struct {
	db domain.IProduct
	service *services.ServiceNotification
}

func NewCreateProduct(db domain.IProduct,service *services.ServiceNotification) *CreateProduct {
	return &CreateProduct{db: db, service: service}
}

func (uc *CreateProduct) Execute(name string, price float32) error {
	product := domain.Product{
		Name:  name,
		Price: price,
	}

	// Guardar el producto
	err := uc.db.Save(product.Name, product.Price)
	if err != nil {
		return err
	}

	// Notificar
	err = uc.service.Execute(product)
	if err != nil {
		return err
	}

	return nil
}

