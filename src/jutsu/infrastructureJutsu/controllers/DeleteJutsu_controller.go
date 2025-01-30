package controllers

import (
	"API-HEXAGONAL/src/jutsu/application/useCaseJutsu"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteJutsuController struct {
	useCase *useCaseJutsu.DeleteJutsu
}

func NewDeleteJutsuController(useCase *useCaseJutsu.DeleteJutsu) *DeleteJutsuController {
	return &DeleteJutsuController{useCase: useCase}
}

func (delete DeleteJutsuController) Run(c *gin.Context) {
	// Extraer el ID de la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al caso de uso para eliminar el jutsu
	rowsAffected, err := delete.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jutsu no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jutsu eliminado correctamente", "rows_affected": rowsAffected})
}
