package controllers

import (
	"API-HEXAGONAL/src/jutsu/application/useCaseJutsu"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateJutsuController struct {
	useCase *useCaseJutsu.CreateJutsu
}

func NewCreateJutsuController(useCase *useCaseJutsu.CreateJutsu) *CreateJutsuController {
	return &CreateJutsuController{useCase: useCase}
}

func (create CreateJutsuController) Run(c *gin.Context) {
	var input struct {
		Name            string `json:"name"`
		JutsuType       string `json:"jutsu_type"`
		Nature          string `json:"nature"`
		DifficultyLevel string `json:"difficulty_level"`
		CreatedBy       string `json:"created_by"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos deben ir completos"})
		return
	}

	rowsAffected, err := create.useCase.Run(input.Name, input.JutsuType, input.Nature, input.DifficultyLevel, input.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Jutsu creado exitosamente", "rows_affected": rowsAffected})
}
