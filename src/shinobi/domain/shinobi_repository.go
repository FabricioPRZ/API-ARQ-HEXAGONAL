package domain

import (
	"API-HEXAGONAL/src/shinobi/domain/entities"
)

type IShinobi interface {
	SaveShinobi(name string, clan string, position string, village string, special_hability string) (int64, error)
	GetAllShinobis() ([]entities.Shinobi, error)
	UpdateShinobi(id int32, name string, clan string, position string, village string, special_hability string) (int64, error)
	DeleteShinobi(id int32) (int64, error)
}
