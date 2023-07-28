package models

import (
	"fmt"
	"log"
	"os"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect() {
	var err error 
	database, err := Postgresql()
	if err != nil {
		log.Fatal("failed to connect database: " , err)
		return 
	}

	DB = database
	// dbSQL, ok := DB.DB()
	// if ok != nil {
	// 	defer dbSQL.Close()
	// }
}

func Postgresql() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,                      // Slow SQL threshold
				LogLevel:      logger.Info,                      // Log level
				Colorful:      true, // Disable color
			},
		),
	})
}

func dsn() string {
	host := "host=127.0.0.1"
	port := "port=5432"
	dbname := "dbname=auth"
	user := "user=postgres"
	password := "password=admin"

	return fmt.Sprintln(host, port, dbname, user, password)
}