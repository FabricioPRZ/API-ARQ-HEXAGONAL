package controllers

import (
	"API-HEXAGONAL/src/jutsu/application/useCaseJutsu"
	"API-HEXAGONAL/src/jutsu/domain/entities"
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
	var input entities.Jutsu

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos deben ir completos"})
		return
	}

	createdJutsu, err := create.useCase.Run(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Jutsu creado exitosamente",
		"jutsu":   createdJutsu,
	})
}
