package postgres

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/golang-migrate/migrate"
	migpg "github.com/golang-migrate/migrate/database/postgres"
)

// MigrateUp applies all up migrations to a pg db.
func MigrateUp(pathToMigs, dsn string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error opening connection for up migrations: ", err)
	}

	driver, err := migpg.WithInstance(db, &migpg.Config{})
	if err != nil {
		log.Fatal("error creating driver for up migrations: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		pathToMigs, // expecting format: "file://../../migrations"
		"postgres", driver)
	if err != nil {
		log.Fatal("error creating up migrations: ", err)
	}

	err = m.Up()
	if err != nil {
		log.Println("error applying up migrations: ", err)
	}

}

// MigrateDown applies all down migrations to a pg db.
func MigrateDown(pathToMigs, dsn string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error opening connection for down migrations: ", err)
	}

	driver, err := migpg.WithInstance(db, &migpg.Config{})
	if err != nil {
		log.Fatal("error creating driver for up migrations: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		pathToMigs, // expecting format: "file://../../migrations"
		"postgres", driver)
	if err != nil {
		log.Fatal("error creating down migrations: ", err)
	}

	err = m.Down()
	if err != nil {
		log.Println("error applying down migrations: ", err)
	}

}

func appendParamsAndArgs(col, val, params, vals string, args []interface{}) (newParams, newVals string, newArgs []interface{}) {
	args = append(args, val)
	if len(args) > 1 {
		params += ", "
		vals += ", "
	}
	params += col
	vals += "$" + strconv.Itoa(len(args))
	return params, vals, args
}
