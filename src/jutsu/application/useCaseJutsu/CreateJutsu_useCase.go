package useCaseJutsu

import (
	"API-HEXAGONAL/src/jutsu/application/repositories"
	"API-HEXAGONAL/src/jutsu/domain"
	"API-HEXAGONAL/src/jutsu/domain/entities"
	"log"
)

type CreateJutsu struct {
	db        domain.IJutsu
	messaging repositories.MessageService
}

func NewCreateJutsu(db domain.IJutsu, messaging repositories.MessageService) *CreateJutsu {
	return &CreateJutsu{db: db, messaging: messaging}
}

func (create *CreateJutsu) Run(jutsu entities.Jutsu) (entities.Jutsu, error) {
	// Guardar el Jutsu en la base de datos
	err := create.db.SaveJutsu(jutsu.Name, jutsu.JutsuType, jutsu.Nature, jutsu.DifficultyLevel, jutsu.CreatedBy)
	if err != nil {
		return entities.Jutsu{}, err
	}

	// Publicar el evento
	err = create.messaging.PublishEvent("Jutsu creado exitosamente", jutsu)
	if err != nil {
		log.Printf("Error al publicar el evento del jutsu: %v", err)
		return entities.Jutsu{}, err
	}

	// Retornar el Jutsu creado
	return jutsu, nil
}
