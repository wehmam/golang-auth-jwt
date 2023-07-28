package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect() {
	var err error
	godotenv.Load(".env")
	database, err := Postgresql()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
		return
	}

	DB = database
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
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		),
	})
}

func dsn() string {
	host := "host=" + os.Getenv("POSTGRES_HOST")
	port := "port=" + os.Getenv("POSTGRES_PORT")
	dbname := "dbname=" + os.Getenv("POSTGRES_DBNAME")
	user := "user=" + os.Getenv("POSTGRES_USER")
	password := "password=" + os.Getenv("POSTGRES_PASSWORD")
	return fmt.Sprintln(host, port, dbname, user, password)
}
