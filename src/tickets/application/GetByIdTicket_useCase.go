package application

import "hexagonal/src/tickets/domain"

type GetByIdTicket struct {
	db domain.ITicket
}


func NewGetByIdTicket(db domain.ITicket)*GetByIdTicket{
	return &GetByIdTicket{db : db}
}

func (getById *GetByIdTicket) Execute(id int32)(domain.Ticket,error){
	return getById.db.GetById(id)
}