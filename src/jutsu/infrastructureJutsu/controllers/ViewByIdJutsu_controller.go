package controllers

import (
	"API-HEXAGONAL/src/jutsu/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetByIdJutsuController(repo domain.IJutsu) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el ID de los parámetros de la URL y convertirlo a int32
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		// Llamar al repositorio para obtener el Jutsu
		jutsu, err := repo.GetJutsuById(int32(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Retornar el jutsu encontrado
		c.JSON(http.StatusOK, jutsu)
	}
}
