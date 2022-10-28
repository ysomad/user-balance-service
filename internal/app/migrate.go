//go:build migrate
// +build migrate

package app

import (
	"os"

	"github.com/ysomad/avito-internship-task/migrate"
)

func init() {
	migrate.Do(migrate.Up, "./migrations", os.Getenv("PG_URL"))
}
