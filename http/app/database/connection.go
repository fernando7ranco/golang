package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	const dbName = "golang"
	const dbUser = "root"
	const dbPassword = ""
	const dbTcp = "@tcp(127.0.0.1:3306)/"
	const config = "charset=utf8&parseTime=True"

	stringConnection := dbUser + ":" + dbPassword + dbTcp + dbName + "?" + config

	db, error := gorm.Open(mysql.Open(stringConnection))
	if error != nil {
		return nil, error
	}

	return db, nil
}
