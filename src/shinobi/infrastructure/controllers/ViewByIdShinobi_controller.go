package controllers

import (
	"API-HEXAGONAL/src/shinobi/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetByIdShinobiController(repo domain.IShinobi) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el ID de los parámetros de la URL y convertirlo a int32
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		// Llamar al repositorio para obtener el Shinobi
		jutsu, err := repo.GetShinobiById(int32(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Retornar el shinobi encontrado
		c.JSON(http.StatusOK, jutsu)
	}
}
