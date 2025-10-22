// Package ramsql contains RamSQL driver registration for xk6-sql.
package bigquery

import (
	"github.com/grafana/xk6-sql/sql"

	// Blank import required for initialization of driver.
	_ "github.com/solcates/go-sql-bigquery"
)

func init() {
	sql.RegisterModule("bigquery")
}
