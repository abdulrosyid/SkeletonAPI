package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER_POSTGRESQL"),
		os.Getenv("DB_PASSWORD_POSTGRESQL"),
		os.Getenv("DB_HOST_POSTGRESQL"),
		os.Getenv("DB_PORT_POSTGRESQL"),
	)
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalf("got an error while initialize database migrations, error: %s", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("got an error while execute database migrations, error: %s", err)
	}
}
