package database

import (
	"Novelas/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

func NewMysql() (*sqlx.DB, error) {

	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		conf.DB.User, conf.DB.Pass, conf.DB.Host, conf.DB.Port, conf.DB.Name)

	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		
		return nil, err
	}

	return db, nil

}
