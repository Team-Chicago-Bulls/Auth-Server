package Rutes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Documentacion(c *gin.Context) {
	enlaceExterno := "https://app.swaggerhub.com/apis-docs/SANTIAGOANDRESDELVAL/Auth-Server/1.0.0"
	c.Redirect(http.StatusFound, enlaceExterno)
}
