package useCaseJutsu

import (
	"API-HEXAGONAL/src/jutsu/domain"
)

type CreateJutsu struct {
	db domain.IJutsu
}

func NewCreateJutsu(db domain.IJutsu) *CreateJutsu {
	return &CreateJutsu{db: db}
}

func (create *CreateJutsu) Run(name string, jutsu_type string, nature string, difficulty_level string, created_by string) (int64, error) {
	return create.db.SaveJutsu(name, jutsu_type, nature, difficulty_level, created_by)
}
