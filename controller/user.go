package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/model"
)

// UserController is the user controller
type UserController struct{}

// GetUser gets the user info
func (ctrl *UserController) GetUser(c *gin.Context) {
	var user model.User

	id := c.Param("id")

	if err := user.GetFirstByID(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
