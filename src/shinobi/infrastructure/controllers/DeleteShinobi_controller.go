package controllers

import (
	"API-HEXAGONAL/src/shinobi/application/useCaseShinobi"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteShinobiController struct {
	useCase *useCaseShinobi.DeleteShinobi
}

func NewDeleteShinobiController(useCase *useCaseShinobi.DeleteShinobi) *DeleteShinobiController {
	return &DeleteShinobiController{useCase: useCase}
}

func (delete DeleteShinobiController) Run(c *gin.Context) {
	// Extraer el ID de la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al caso de uso para eliminar el shinobi
	rowsAffected, err := delete.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shinobi no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shinobi eliminado correctamente", "rows_affected": rowsAffected})
}
