package database

import (
	"fmt"
	"go-boilerplate/shared/config"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
)

func OpenMysqlConn() (*gorm.DB, error) {
	log.Info("Start open mysql connection...")
	connectionString := GetConnectionString()

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(32)
	sqlDB.SetMaxIdleConns(64)
	sqlDB.SetConnMaxLifetime(30 * time.Second)
	return db, nil
}

func GetConnectionString() string {
	conf := config.GetConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		conf.GetDatabaseMySQLUsername(),
		conf.GetDatabaseMySQLPassword(),
		conf.GetDatabaseMySQLHost(),
		"3306",
		conf.GetDatabaseMySQLDBName(),
	)
}
