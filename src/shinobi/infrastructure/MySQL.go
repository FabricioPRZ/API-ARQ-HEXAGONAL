package infrastructure

import (
	"API-HEXAGONAL/src/core"
	"API-HEXAGONAL/src/shinobi/domain"
	"API-HEXAGONAL/src/shinobi/domain/entities"
)

// MySQLRepository implementa la interfaz IShinobi
type MySQLRepository struct{}

// NewMySQLRepository crea una nueva instancia del repositorio
func NewMySQLRepositoryShinobi() domain.IShinobi {
	return &MySQLRepository{}
}

// Save almacena un nuevo shinobi en la base de datos
func (repo *MySQLRepository) SaveShinobi(name string, clan string, position string, village string, specialHability string) (int64, error) {
	shinobi := entities.Shinobi{Name: name, Clan: clan, Position: position, Village: village, SpecialHability: specialHability}
	result := core.DB.Create(&shinobi)
	return result.RowsAffected, result.Error
}

// GetAll recupera todos los shinobis de la base de datos
func (repo *MySQLRepository) GetAllShinobis() ([]entities.Shinobi, error) {
	var shinobis []entities.Shinobi
	result := core.DB.Find(&shinobis)
	return shinobis, result.Error
}

// Update actualiza un shinobi en la base de datos
func (repo *MySQLRepository) UpdateShinobi(id int32, name string, clan string, position string, village string, specialHability string) (int64, error) {
	result := core.DB.Model(&entities.Shinobi{}).Where("id_shinobi = ?", id).Updates(entities.Shinobi{
		Name: name, Clan: clan, Position: position, Village: village, SpecialHability: specialHability,
	})
	return result.RowsAffected, result.Error
}

// Delete elimina un shinobi de la base de datos
func (repo *MySQLRepository) DeleteShinobi(id int32) (int64, error) {
	result := core.DB.Where("id_shinobi = ?", id).Delete(&entities.Shinobi{})
	return result.RowsAffected, result.Error
}
