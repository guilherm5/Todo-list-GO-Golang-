package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud-gingonic-mvc/autentication"
)

func Autentication(c *gin.Engine) {
	c.GET("/Credentials", autentication.Credentials())
	c.GET("/Token", autentication.Token())

}
