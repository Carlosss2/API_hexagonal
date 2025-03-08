package repositories

import "hexagonal/src/products/domain"

type INotification interface {
	PublishEvent(event string, product domain.Product) error}