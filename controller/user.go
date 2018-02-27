package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/middleware"
	"github.com/hyperjiang/gin-skeleton/model"
)

// UserController is the user controller
type UserController struct{}

// Login struct
type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

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

// Login a user
func (ctrl *UserController) Login(c *gin.Context) {
	type LoginResult struct {
		model.User
		Token  string `json:"token"`
		Expire string `json:"expire"`
	}
	var form Login
	if err := c.ShouldBind(&form); err == nil {
		user, err2 := model.LoginByEmailAndPassword(form.Email, form.Password)
		if err2 != nil {
			c.JSON(http.StatusOK, gin.H{"error": err2.Error()})
		} else {
			token, expire, err3 := middleware.Auth().TokenGenerator(user.Name)
			if err3 != nil {
				c.JSON(http.StatusOK, gin.H{"error": err3.Error()})
			}
			res := LoginResult{
				User:   user,
				Token:  token,
				Expire: expire.Format(time.RFC3339),
			}
			c.JSON(http.StatusOK, res)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
