package controlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud-gingonic-mvc/database"
	"github.com/guilherm5/crud-gingonic-mvc/models"
)

func CreateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var H = database.Init()
		var creating models.Todo_List

		err := c.ShouldBindJSON(&creating)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error in ShouldBind ": err.Error(),
			})
			return
		}

		err = H.Create(&creating).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error in Creating ": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Created the Todo ": creating,
		})
	}
}

func ListTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var H = database.Init()
		var Listing []models.Todo_List

		err := H.Find(&Listing).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error in listing todos ": err.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"todos ": Listing,
		})
	}
}

func DeleteTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		var H = database.Init()
		var request models.Todo_List_DELL

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//Begin permite que façamos uma serie de operações em nosso BD
		tx := H.Begin()
		for _, id := range request.Ids {
			var todo models.Todo_List

			if result := tx.First(&todo, id); result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": result.Error,
				})
				return
			}

			if err := tx.Delete(&todo).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				return
			}
		}
		//Commit confirmas as operações que realizei
		tx.Commit()

		c.JSON(http.StatusOK, gin.H{
			"message ": request,
		})
	}
}
func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var H = database.Init()
		var updating models.Todo_List

		if err := c.ShouldBindJSON(&updating); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error in ShouldBind": err.Error(),
			})
			return
		}

		var todo models.Todo_List
		if result := H.Where("id = ?", updating.ID).First(&todo); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error id capture ": result.Error,
			})
			return
		}

		todo.Tarefa = updating.Tarefa
		todo.ResumoTarefa = updating.ResumoTarefa
		todo.Realizado = updating.Realizado

		if err := H.Save(&todo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error in update todo ": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Todo atualizado com sucesso": todo,
		})
	}
}
