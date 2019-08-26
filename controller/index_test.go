package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/config"
	"github.com/hyperjiang/php"
	"github.com/stretchr/testify/suite"
)

type IndexTestSuite struct {
	suite.Suite
	rec     *httptest.ResponseRecorder
	context *gin.Context
	app     *gin.Engine
	ctrl    *IndexController
}

func (suite *IndexTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	config.Server.Mode = gin.ReleaseMode

	suite.rec = httptest.NewRecorder()
	suite.context, suite.app = gin.CreateTestContext(suite.rec)
	suite.app.LoadHTMLGlob("../view/*")
	suite.ctrl = new(IndexController)
}

func (suite *IndexTestSuite) TestGetIndex() {
	suite.ctrl.GetIndex(suite.context)
	suite.Equal(200, suite.rec.Code)
}

func (suite *IndexTestSuite) TestGetVersion() {
	suite.ctrl.GetVersion(suite.context)
	suite.Equal(200, suite.rec.Code)
	suite.Equal("{\"version\":\"v0.1\"}", php.Trim(suite.rec.Body.String()))
}

func TestIndexTestSuite(t *testing.T) {
	suite.Run(t, new(IndexTestSuite))
}
