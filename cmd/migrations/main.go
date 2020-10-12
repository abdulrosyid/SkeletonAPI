package main

import (
	"fmt"
	configEnv "github.com/joho/godotenv"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func main() {
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER_POSTGRESQL"),
		os.Getenv("DB_PASSWORD_POSTGRESQL"),
		os.Getenv("DB_HOST_POSTGRESQL"),
		os.Getenv("DB_PORT_POSTGRESQL"),
		os.Getenv("DB_NAME_POSTGRESQL"),
	)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalf("got an error while initialize database migrations, error: %s", err)
	}

	m.Up()
}
