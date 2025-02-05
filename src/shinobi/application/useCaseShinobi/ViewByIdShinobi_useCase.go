package useCaseShinobi

import (
	"API-HEXAGONAL/src/shinobi/domain"
	"API-HEXAGONAL/src/shinobi/domain/entities"
)

type ViewShinobiById struct {
	db domain.IShinobi
}

func NewViewShinobiById(db domain.IShinobi) *ViewShinobiById {
	return &ViewShinobiById{db: db}
}

func (view *ViewShinobiById) Run(id int32) (entities.Shinobi, error) {
	return view.db.GetShinobiById(id)
}
