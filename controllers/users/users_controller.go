package user_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mcmuralishclint/bookstore-users-api/domain/users"
	"github.com/mcmuralishclint/bookstore-users-api/services"
	"github.com/mcmuralishclint/bookstore-users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.CreateUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, savgetErrErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me")
// }
