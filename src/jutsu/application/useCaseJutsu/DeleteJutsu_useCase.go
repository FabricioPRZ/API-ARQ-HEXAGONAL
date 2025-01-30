package useCaseJutsu

import (
	"API-HEXAGONAL/src/jutsu/domain"
)

type DeleteJutsu struct {
	db domain.IJutsu
}

func NewDeleteJutsu(db domain.IJutsu) *DeleteJutsu {
	return &DeleteJutsu{db: db}
}

func (delete *DeleteJutsu) Run(id int32) (int64, error) {
	return delete.db.DeleteJutsu(id)
}
