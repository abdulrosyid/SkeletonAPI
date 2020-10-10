package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func CreateDBConnection(descriptor string, maxIdle int, MaxOpen int) *sql.DB {
	db, err := sql.Open("postgres", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(MaxOpen)
	db.SetConnMaxLifetime(time.Minute * 3)

	err = db.Ping()
	if err != nil {
		log.Info("not connect database")
		log.Fatal(err)
	}

	return db
}

func PostgreSqlDB() *sql.DB {
	maxIdle, _:= strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen, _:= strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	return CreateDBConnection(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		os.Getenv("DB_HOST_POSTGRESQL"), os.Getenv("DB_PORT_POSTGRESQL"), 
		os.Getenv("DB_USER_POSTGRESQL"), os.Getenv("DB_PASSWORD_POSTGRESQL"), 
		os.Getenv("DB_NAME_POSTGRESQL")),
		maxIdle, maxOpen)
}


