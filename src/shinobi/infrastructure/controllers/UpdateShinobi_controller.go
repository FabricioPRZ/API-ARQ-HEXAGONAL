package controllers

import (
	"API-HEXAGONAL/src/shinobi/application/useCaseShinobi"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateShinobiController struct {
	useCase *useCaseShinobi.UpdateShinobi
}

func NewUpdateShinobiController(useCase *useCaseShinobi.UpdateShinobi) *UpdateShinobiController {
	return &UpdateShinobiController{useCase: useCase}
}

func (update UpdateShinobiController) Run(c *gin.Context) {
	var input struct {
		Name            string `json:"name"`
		Clan            string `json:"clan"`
		Position        string `json:"position"`
		Village         string `json:"village"`
		SpecialHability string `json:"special_hability"`
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

	rowsAffected, err := update.useCase.Run(int32(id), input.Name, input.Clan, input.Position, input.Village, input.SpecialHability)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shinobi actualizado correctamente", "rows_affected": rowsAffected})
}
