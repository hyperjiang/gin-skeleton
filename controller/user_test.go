package controller

import (
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

func (suite *UserTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	config.Server.Mode = gin.ReleaseMode

	suite.rec = httptest.NewRecorder()
	suite.context, suite.app = gin.CreateTestContext(suite.rec)
	suite.ctrl = new(UserController)
}

func (suite *UserTestSuite) TestGetUser() {
	suite.context.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	suite.ctrl.GetUser(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
