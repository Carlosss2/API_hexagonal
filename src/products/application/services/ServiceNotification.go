package services

import (
	"hexagonal/src/products/application/repositories"
	"hexagonal/src/products/domain"
	"log"
)

type ServiceNotification struct {
	notification repositories.INotification
}

func NewServiceNotification(notification repositories.INotification) *ServiceNotification {
	return &ServiceNotification{notification: notification}
}

func (sv *ServiceNotification) Execute(product domain.Product) error {
	log.Println("Notificando nuevo platillo")

	err := sv.notification.PublishEvent("producto creado", product)
	if err != nil {
		log.Printf("Error al publicar: %v", err)
		return err
	}
	return nil
}
