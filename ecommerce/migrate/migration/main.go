package main

import (
	"log"
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/joshua468/ecommerce/config"

	"github.com/joshua468/ecommerce/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                   config.Envs.DBUser,
		Password:               config.Envs.DBPassword,
		Addr:                   config.Envs.DBAddr,
		DbName:                 config.Envs.DBName,
		Net:                    "tcp",
		AllowedNativePasswords: "true",
		ParseTime:              "True",
	})
	if err != nil {
		log.Fatal(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
