package domain

import "API-HEXAGONAL/src/jutsu/domain/entities"

type IJutsu interface {
	SaveJutsu(name string, jutsu_type string, nature string, difficulty_level string, created_by string) (int64, error)
	GetAllJutsus() ([]entities.Jutsu, error)
	GetJutsuById(id int32) (entities.Jutsu, error)
	UpdateJutsu(id int32, name string, jutsu_type string, nature string, difficulty_level string, created_by string) (int64, error)
	DeleteJutsu(id int32) (int64, error)
}
