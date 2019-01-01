package integration

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	host := os.Getenv("DB_HOST")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")

	var err error

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/drop_rules", user, pass, host, port))

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func TestFirst(t *testing.T) {
	db.Exec("hi")
	t.Fatalf("bad")
}
