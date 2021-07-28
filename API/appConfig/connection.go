package appConfig

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "sqlserver://sa:1q2w3eR@dev01.softwareworx.co.za?"
	if database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		panic("Error connecting to database")
	} else {
		sqlDB, err := database.DB()
		if err != nil {
			fmt.Println("errror", err)
			return
		}
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
		DB = database
	}

}
