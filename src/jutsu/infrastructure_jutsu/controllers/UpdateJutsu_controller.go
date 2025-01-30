package controllers

import (
	"API-HEXAGONAL/src/jutsu/application/useCaseJutsu"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateJutsuController struct {
	useCase *useCaseJutsu.UpdateJutsu
}

func NewUpdateJutsuController(useCase *useCaseJutsu.UpdateJutsu) *UpdateJutsuController {
	return &UpdateJutsuController{useCase: useCase}
}

func (update UpdateJutsuController) Run(c *gin.Context) {
	var input struct {
		Name            string `json:"name"`
		JutsuType       string `json:"jutsu_type"`
		Nature          string `json:"nature"`
		DifficultyLevel string `json:"difficulty_level"`
		CreatedBy       string `json:"created_by"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos incorrectos"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	rowsAffected, err := update.useCase.Run(int32(id), input.Name, input.JutsuType, input.Nature, input.DifficultyLevel, input.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jutsu actualizado correctamente", "rows_affected": rowsAffected})
}
