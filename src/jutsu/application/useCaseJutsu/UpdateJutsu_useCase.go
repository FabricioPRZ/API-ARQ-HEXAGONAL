package useCaseJutsu

import (
	"API-HEXAGONAL/src/jutsu/domain"
)

type UpdateJutsu struct {
	db domain.IJutsu
}

func NewUpdateJutsu(db domain.IJutsu) *UpdateJutsu {
	return &UpdateJutsu{db: db}
}

func (update *UpdateJutsu) Run(id int32, name string, jutsu_type string, nature string, difficulty_level string, created_by string) (int64, error) {
	return update.db.UpdateJutsu(id, name, jutsu_type, nature, difficulty_level, created_by)
}
