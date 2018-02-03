package model

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/hyperjiang/gin-skeleton/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
)

// DBInstance is a singleton DB instance
type DBInstance struct {
	initializer func() interface{}
	instance    interface{}
	once        sync.Once
}

var (
	dbInstance *DBInstance
)

// Instance gets the singleton instance
func (i *DBInstance) Instance() interface{} {
	i.once.Do(func() {
		i.instance = i.initializer()
	})
	return i.instance
}

func dbInit() interface{} {
	db, err := gorm.Open(config.Database.Dialect, config.Database.DSN)
	if err != nil {
		glog.Fatalf("Cannot connect to database: %v", err)
	}

	// sql log
	if config.Server.Mode != gin.ReleaseMode {
		db.LogMode(true)
	}

	stdDB := db.DB()
	stdDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	stdDB.SetMaxOpenConns(config.Database.MaxOpenConns)

	return db
}

// DB returns the database instance
func DB() *gorm.DB {
	return dbInstance.Instance().(*gorm.DB)
}

func init() {
	dbInstance = &DBInstance{initializer: dbInit}
}
