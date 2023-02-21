package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igroks/sd-project/backend/app/config"
	"github.com/igroks/sd-project/backend/app/database"
	"github.com/igroks/sd-project/backend/app/models"
)

func Add(c *gin.Context) {
	var requestMsg models.Request
	var db *sql.DB
	var name string
	databaseName := c.Param("databaseName")

	if err := c.ShouldBindJSON(&requestMsg); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	db = database.OpenConn(config.GetDatabaseConfig(databaseName))

	sqlQuery := `INSERT INTO items (name) VALUES ($1) RETURNING name`
	err := db.QueryRow(sqlQuery, *requestMsg.Item).Scan(&name)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	defer db.Close()

	responseMsg := models.Response{
		Message: "New item added: " + name,
	}

	c.JSON(http.StatusOK, &responseMsg)
}

func List(c *gin.Context) {
	var db *sql.DB
	databaseName := c.Param("databaseName")

	db = database.OpenConn(config.GetDatabaseConfig(databaseName))

	sqlQuery := `SELECT name FROM items`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	var items []string

	for rows.Next() {
		var item string
		rows.Scan(&item)
		items = append(items, item)
	}

	defer rows.Close()
	defer db.Close()

	c.JSON(http.StatusOK, map[string][]string{
		"items": items,
	})
}
