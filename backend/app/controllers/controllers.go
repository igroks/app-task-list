package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igroks/sd-project/backend/app/config"
	"github.com/igroks/sd-project/backend/app/database"
	"github.com/igroks/sd-project/backend/app/models"
)

func Add(c *gin.Context) {
	var requestItem models.Item
	databaseName := c.Param("databaseName")

	if err := c.ShouldBindJSON(&requestItem); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	db := database.OpenConn(config.Env.Database[databaseName])

	sqlQuery := `INSERT INTO items (name, createdAt, duplicated) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlQuery, requestItem.Name, requestItem.CreatedAt, requestItem.Duplicated)

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
	databaseName := c.Param("databaseName")

	db := database.OpenConn(config.Env.Database[databaseName])

	sqlQuery := `SELECT id, name, createdAt, duplicated FROM items`
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
		rows.Scan(&item.Id, &item.Name, &item.CreatedAt, &item.Duplicated)
		items = append(items, item)
	}

	defer rows.Close()
	defer db.Close()

	c.JSON(http.StatusOK, map[string][]models.Item{
		"items": items,
	})
}

func Delete(c *gin.Context) {
	var requestItem models.DeleteItemsRequest
	databaseName := c.Param("databaseName")

	if err := c.ShouldBindJSON(&requestItem); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	db := database.OpenConn(config.Env.Database[databaseName])

	sqlQuery := `DELETE FROM items WHERE id = $1`

	_, err := db.Exec(sqlQuery, requestItem.Id)

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
