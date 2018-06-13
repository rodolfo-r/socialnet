package storage_test

import (
	"flag"
	"os"

	"github.com/techmexdev/socialnet/pkg/server/storage/postgres"
)

var testMemo bool
var dsn string

func init() {
	flag.BoolVar(&testMemo, "memo", false, "tests the memo package")
	flag.Parse()

	if !testMemo {
		dsn = os.Getenv("PG_TEST_DSN")

		postgres.MigrateDown("file://postgres/migrations", dsn)
		postgres.MigrateUp("file://postgres/migrations", dsn)
	}
}
