package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseBad ...
func ResponseBad(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}

// ResponseOK ...
func ResponseOK(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": response})
}
