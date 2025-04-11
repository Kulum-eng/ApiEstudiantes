package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"API-HEXAGONAL/src/materia/domain"
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
