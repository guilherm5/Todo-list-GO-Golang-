package controlers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud-gingonic-mvc/controlers"
)

func TodoList(c *gin.Engine) {
	c.GET("/Todo", controlers.ListTodo())
	c.POST("/Todo", controlers.CreateTodo())
	c.DELETE("/Todo", controlers.DeleteTodos())
	c.PUT("/Todo", controlers.UpdateTodo())
}
