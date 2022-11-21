package model

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBInstance is a singleton DB instance
type DBInstance struct {
	initializer func() any
	instance    any
	once        sync.Once
}

var (
	dbInstance *DBInstance
)

// Instance gets the singleton instance
func (i *DBInstance) Instance() any {
	i.once.Do(func() {
		i.instance = i.initializer()
	})
	return i.instance
}

func dbInit() any {
	lv := logger.Error
	if config.Server.Mode != gin.ReleaseMode {
		lv = logger.Info // output debug logs in dev mode
	}

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(lv),
	}

	db, err := gorm.Open(mysql.Open(config.Database.DSN), cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to database")
	}

	stdDB, _ := db.DB()
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
