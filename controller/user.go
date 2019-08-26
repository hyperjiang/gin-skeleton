package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/model"
)

// UserController is the user controller
type UserController struct{}

// Signup struct
type Signup struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Password2 string `form:"password2" json:"password2" binding:"required"`
}

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

// SignupForm shows the signup form
func (ctrl *UserController) SignupForm(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// LoginForm shows the login form
func (ctrl *UserController) LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// Signup a new user
func (ctrl *UserController) Signup(c *gin.Context) {
	var form Signup
	if err := c.ShouldBind(&form); err == nil {

		if form.Password != form.Password2 {
			c.JSON(http.StatusOK, gin.H{"error": "Password does not match with conform password"})
			return
		}

		var user model.User

		user.Name = form.Name
		user.Email = form.Email
		user.Password = form.Password

		if err := user.Signup(); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
