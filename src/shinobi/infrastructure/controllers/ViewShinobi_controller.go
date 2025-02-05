package controllers

import (
	"API-HEXAGONAL/src/shinobi/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllShinobisController(repo domain.IShinobi) gin.HandlerFunc {
	return func(c *gin.Context) {
		shinobis, err := repo.GetAllShinobis()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, shinobis)
	}
}
