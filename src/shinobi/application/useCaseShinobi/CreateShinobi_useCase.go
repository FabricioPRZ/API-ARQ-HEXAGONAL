package useCaseShinobi

import (
	"API-HEXAGONAL/src/shinobi/domain"
)

type CreateShinobi struct {
	db domain.IShinobi
}

func NewCreateShinobi(db domain.IShinobi) *CreateShinobi {
	return &CreateShinobi{db: db}
}

func (create *CreateShinobi) Run(name string, clan string, position string, village string, special_hability string) (int64, error) {
	return create.db.SaveShinobi(name, clan, position, village, special_hability)
}
