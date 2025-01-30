package useCaseShinobi

import (
	"API-HEXAGONAL/src/shinobi/domain"
	"API-HEXAGONAL/src/shinobi/domain/entities"
)

type ViewShinobis struct {
	db domain.IShinobi
}

func NewViewShinobis(db domain.IShinobi) *ViewShinobis {
	return &ViewShinobis{db: db}
}

func (view *ViewShinobis) Run() ([]entities.Shinobi, error) {
	return view.db.GetAllShinobis()
}
