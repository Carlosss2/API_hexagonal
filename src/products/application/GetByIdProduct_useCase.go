package application

import "hexagonal/src/products/domain"

type GetByIdProduct struct {
	db domain.IProduct
}

func NewGetByIdProduct(db domain.IProduct)*GetByIdProduct{
	return &GetByIdProduct{db: db}
}

func (getById *GetByIdProduct) Execute(id int32)(domain.Product,error){
	return getById.db.GetById(id)
}