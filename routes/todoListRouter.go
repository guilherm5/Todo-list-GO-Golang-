package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud-gingonic-mvc/autentication"
	"github.com/guilherm5/crud-gingonic-mvc/controlers"
)

func TodoList(c *gin.Engine) {
	c.GET("/Todo", autentication.MiddlewareAutentication(), controlers.ListTodo())
	c.POST("/Todo", autentication.MiddlewareAutentication(), controlers.CreateTodo())
	c.DELETE("/Todo", controlers.DeleteTodos())
	c.PUT("/Todo", controlers.UpdateTodo())
}
