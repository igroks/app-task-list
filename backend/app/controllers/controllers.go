package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igroks/app-task-list/backend/app/config"
	"github.com/igroks/app-task-list/backend/app/database"
	"github.com/igroks/app-task-list/backend/app/models"
)

func Add(c *gin.Context) {
	var requestItem models.Item

	if err := c.ShouldBindJSON(&requestItem); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	db := database.OpenConn(config.Env.Database)

	sqlQuery := `INSERT INTO items (name, createdAt) VALUES ($1, $2)`
	_, err := db.Exec(sqlQuery, requestItem.Name, requestItem.CreatedAt)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	defer db.Close()

	c.JSON(http.StatusOK, models.Response{
		Message: "New item added",
	})
}

func List(c *gin.Context) {
	var items []models.Item

	db := database.OpenConn(config.Env.Database)

	sqlQuery := `SELECT id, name, createdAt FROM items`
	rows, err := db.Query(sqlQuery)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	for rows.Next() {
		var item models.Item
		rows.Scan(&item.Id, &item.Name, &item.CreatedAt)
		items = append(items, item)
	}

	defer rows.Close()
	defer db.Close()

	c.JSON(http.StatusOK, map[string][]models.Item{
		"items": items,
	})
}

func Delete(c *gin.Context) {
	taskId := c.Param("taskId")

	db := database.OpenConn(config.Env.Database)

	sqlQuery := `DELETE FROM items WHERE id = $1`

	_, err := db.Exec(sqlQuery, taskId)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Item deleted successfully",
	})
}
