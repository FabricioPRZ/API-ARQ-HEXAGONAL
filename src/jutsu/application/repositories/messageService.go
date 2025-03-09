package repositories

import "API-HEXAGONAL/src/jutsu/domain/entities"

type MessageService interface {
	PublishEvent(evenType string, jutsu entities.Jutsu) error
}
