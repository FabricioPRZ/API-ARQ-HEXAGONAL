package controllers

import (
	"API-HEXAGONAL/src/shinobi/application/useCaseShinobi"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateShinobiController struct {
	useCase *useCaseShinobi.CreateShinobi
}

func NewCreateShinobiController(useCase *useCaseShinobi.CreateShinobi) *CreateShinobiController {
	return &CreateShinobiController{useCase: useCase}
}

func (create CreateShinobiController) Run(c *gin.Context) {
	var input struct {
		Name            string `json:"name"`
		Clan            string `json:"clan"`
		Position        string `json:"position"`
		Village         string `json:"village"`
		SpecialHability string `json:"special_hability"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos deben ir completos"})
		return
	}

	rowsAffected, err := create.useCase.Run(input.Name, input.Clan, input.Position, input.Village, input.SpecialHability)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Shinobi creado exitosamente", "rows_affected": rowsAffected})
}
