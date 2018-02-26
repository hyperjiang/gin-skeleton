package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/config"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	rec     *httptest.ResponseRecorder
	context *gin.Context
	app     *gin.Engine
	ctrl    *UserController
}

var (
	name     = "hyper"
	email    = "hyperjiang@gmail.com"
	password = "123456"
)

func (suite *UserTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	config.Server.Mode = gin.ReleaseMode

	suite.rec = httptest.NewRecorder()
	suite.context, suite.app = gin.CreateTestContext(suite.rec)
	suite.app.LoadHTMLGlob("../view/*")
	suite.ctrl = new(UserController)
}

func (suite *UserTestSuite) TestGetUser() {
	suite.context.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	suite.ctrl.GetUser(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func (suite *UserTestSuite) TestSignupForm() {
	suite.ctrl.SignupForm(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func (suite *UserTestSuite) TestLoginForm() {
	suite.ctrl.LoginForm(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func (suite *UserTestSuite) TestSignup() {
	url := fmt.Sprintf(
		"/signup?name=%s&email=%s&password=%s&password2=%s",
		name,
		email,
		password,
		password,
	)
	suite.context.Request, _ = http.NewRequest(
		"POST",
		url,
		bytes.NewBufferString(""),
	)

	suite.ctrl.Signup(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func (suite *UserTestSuite) TestLogin() {
	url := fmt.Sprintf(
		"/login?email=%s&password=%s",
		email,
		password,
	)
	suite.context.Request, _ = http.NewRequest(
		"POST",
		url,
		bytes.NewBufferString(""),
	)

	suite.ctrl.Login(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
