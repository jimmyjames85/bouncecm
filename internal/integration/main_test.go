package integration

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"testing"
)

var Database *sql.DB

func TestMain(m *testing.M) {
	host := os.Getenv("DB_HOST")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	var err error

	Database, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName))
	log.Println("the Server has Started")

	if err != nil {
		panic(err)
	}

	if err = Database.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database has Started")

	e := m.Run()

	Database.Close() // Needed?

	os.Exit(e)
}

func peekResponse(response *http.Response) {
	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(responseDump))
}

func peekRequest(request *http.Request) {
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))
}
