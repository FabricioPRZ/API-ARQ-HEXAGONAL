package useCaseJutsu

import (
	"API-HEXAGONAL/src/jutsu/domain"
	"API-HEXAGONAL/src/jutsu/domain/entities"
)

type ViewJutsu struct {
	db domain.IJutsu
}

func NewViewJutsu(db domain.IJutsu) *ViewJutsu {
	return &ViewJutsu{db: db}
}

func (view *ViewJutsu) Run() ([]entities.Jutsu, error) {
	return view.db.GetAllJutsus()
}
