package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igroks/sd-project/backend/app/models"
)

func Add(c *gin.Context) {
	var requestMsg models.MatchRequest

	if err := c.ShouldBindJSON(&requestMsg); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	responseMsg := models.MatchResponse{
		Message: "Item a ser inserido:" + *requestMsg.Item,
	}

	c.JSON(http.StatusOK, &responseMsg)
}
