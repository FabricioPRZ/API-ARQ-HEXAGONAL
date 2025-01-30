package infrastructure_jutsu

import (
	"API-HEXAGONAL/src/core"
	"API-HEXAGONAL/src/jutsu/domain"
	"API-HEXAGONAL/src/jutsu/domain/entities"
)

// MySQLRepository implementa la interfaz IJutsu
type MySQLRepository struct{}

// NewMySQLRepository crea una nueva instancia del repositorio
func NewMySQLRepositoryJutsu() domain.IJutsu {
	return &MySQLRepository{}
}

// Save almacena un nuevo jutsu en la base de datos
func (repo *MySQLRepository) SaveJutsu(name string, jutsu_type string, nature string, difficulty_level string, created_by string) (int64, error) {
	jutsu := entities.Jutsu{Name: name, JutsuType: jutsu_type, Nature: nature, DifficultyLevel: difficulty_level, CreatedBy: created_by}
	result := core.DB.Create(&jutsu)
	return result.RowsAffected, result.Error
}

// GetAll recupera todos los jutsus de la base de datos
func (repo *MySQLRepository) GetAllJutsus() ([]entities.Jutsu, error) {
	var jutsus []entities.Jutsu
	result := core.DB.Find(&jutsus)
	return jutsus, result.Error
}

// Update actualiza un jutsu en la base de datos
func (repo *MySQLRepository) UpdateJutsu(id int32, name string, jutsu_type string, nature string, difficulty_level string, created_by string) (int64, error) {
	result := core.DB.Model(&entities.Jutsu{}).Where("id_jutsu = ?", id).Updates(entities.Jutsu{
		Name: name, JutsuType: jutsu_type, Nature: nature, DifficultyLevel: difficulty_level, CreatedBy: created_by,
	})
	return result.RowsAffected, result.Error
}

// Delete elimina un jutsu de la base de datos
func (repo *MySQLRepository) DeleteJutsu(id int32) (int64, error) {
	result := core.DB.Where("id_jutsu = ?", id).Delete(&entities.Jutsu{})
	return result.RowsAffected, result.Error
}
