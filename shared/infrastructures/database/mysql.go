package database

import (
	"fmt"
	"go-boilerplate/shared/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMysqlConn(config config.Database) (*gorm.DB, error) {
	fmt.Println("Start open mysql connection...")
	connectionString := GetConnectionString(config)

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

func GetConnectionString(conf config.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=true&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
}
