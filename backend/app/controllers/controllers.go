package controllers

import (
	"database/sql"
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

	if err := c.ShouldBindJSON(&requestMsg); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	switch databaseName := c.Param("databaseName"); databaseName {
	case "db1":
		db = database.OpenConn(config.GetDb1())
	case "db2":
		db = database.OpenConn(config.GetDb2())
	default:
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid database parameter",
		})
		return
	}

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
