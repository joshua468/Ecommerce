package db

import (
	"database/sql"

	"gorm.io/driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB ,error) {
db,err := sql.Open("mysql",cfg.FormatDSN)
if err!= nil {
	log.Fatal(err)
}
return db,nil
}
