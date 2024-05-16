package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/joshua468/ecommerce/cmd/api"
	"github.com/joshua468/ecommerce/cmd/api/db"
	"gorm.io/driver/mysql"
)

func main() {

	db,err:= db.NewMySQLStorage(mysql.Config {
		User:				config.Envs.DBUser,
		Password:			config.Envs.DBPassword,
		Addr:				config.Envs.DBAddr,
		DbName:				config.Envs.DBName,
		Net:"tcp",
		AllowedNativePasswords:"true",
		ParseTime:"True",
	})
	if err!= nil {
		log.Fatal(err)
	}
	}
	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

	func InitStorage(db *sql.DB) {
		err:= db.Ping()
		if err!= nil {
			log.Fatal(err)
		}
		log.Println("DB is successfully connected")
	}



