package repositories

import (
	"API-HEXAGONAL/src/jutsu/domain/entities"
	"log"
)

type ServiceNotification struct {
	imageService MessageService
}

func NewServiceNotification(imageService MessageService) *ServiceNotification {
	return &ServiceNotification{imageService: imageService}
}

func (sn *ServiceNotification) Notify(jutsu entities.Jutsu) error {
	log.Println("Se está creando un jutsu")

	err := sn.imageService.PublishEvent("Jutsu Creado correctamente", jutsu)
	if err != nil {
		log.Println("Ocurrió un error al enviar el mensaje", err)
		return err
	}
	return nil
}
