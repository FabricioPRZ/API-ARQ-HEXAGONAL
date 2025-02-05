package controllers

import (
	"API-HEXAGONAL/src/jutsu/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllJutsuController(repo domain.IJutsu) gin.HandlerFunc {
	return func(c *gin.Context) {
		jutsus, err := repo.GetAllJutsus()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, jutsus)
	}
}
