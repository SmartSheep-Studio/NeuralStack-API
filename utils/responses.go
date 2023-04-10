package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func AbortWithDatasourceError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithError(http.StatusNotFound, err)
	} else if errors.Is(err, gorm.ErrInvalidValue) || errors.Is(err, gorm.ErrDuplicatedKey) {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
