package useCaseShinobi

import (
	"API-HEXAGONAL/src/shinobi/domain"
)

type UpdateShinobi struct {
	db domain.IShinobi
}

func NewUpdateShinobi(db domain.IShinobi) *UpdateShinobi {
	return &UpdateShinobi{db: db}
}

func (update *UpdateShinobi) Run(id int32, name string, clan string, position string, village string, special_hability string) (int64, error) {
	return update.db.UpdateShinobi(id, name, clan, position, village, special_hability)
}
