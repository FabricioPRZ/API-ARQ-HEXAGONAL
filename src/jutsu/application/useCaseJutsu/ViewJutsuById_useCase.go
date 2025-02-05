package useCaseJutsu

import (
	"API-HEXAGONAL/src/jutsu/domain"
	"API-HEXAGONAL/src/jutsu/domain/entities"
)

type ViewJutsuById struct {
	db domain.IJutsu
}

func NewViewJutsuById(db domain.IJutsu) *ViewJutsuById {
	return &ViewJutsuById{db: db}
}

func (view *ViewJutsuById) Run(id int32) (entities.Jutsu, error) {
	return view.db.GetJutsuById(id)
}
