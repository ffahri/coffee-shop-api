package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"os"
)

func Start() *sqlx.DB {
	pass := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USERNAME")
	db, err := sqlx.Open("mysql", username+":"+pass+"@/COFFEESHOP")
	if err != nil {
		panic("Could not open database connection " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("Could not open database connection " + err.Error())
	}
	log.Info().Msg("Connection established")
	return db
}
