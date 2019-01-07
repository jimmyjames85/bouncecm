package integration

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseInsert(t *testing.T) {
	_, err := Database.Exec(`SELECT * From bounce_rule WHERE id = 174`)

	if err != nil {
		t.Errorf("SELECT query failed")
	}
}
