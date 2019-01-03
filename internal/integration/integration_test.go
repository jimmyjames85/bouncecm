package integration

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jimmyjames85/bouncecm/internal/config"
	"github.com/jimmyjames85/bouncecm/internal/sgbouncewizard"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func init() {
	host := os.Getenv("DB_HOST")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")

	var err error

	database, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/drop_rules", user, pass, host, port))

	if err != nil {
		panic(err)
	}

	if err = database.Ping(); err != nil {
		panic(err)
	}
}

func TestFirst(t *testing.T) {
	database.Exec("hi")
}

func TestGetBounceRules(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/bounce_rules")
	log.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
func TestGetChangelogs(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/changelogs")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestChangeLogsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/changelogs", nil)

	if err != nil {
		t.Fatal(err)
	}
	cfg, err := config.LoadConfig()
	rr := httptest.NewRecorder()
	srv, err := sgbouncewizard.NewServer(cfg)
	handler := http.HandlerFunc(srv.GetChangelog)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestBounceRuleHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/bounce_rules", nil)

	if err != nil {
		t.Fatal(err)
	}

	cfg, err := config.LoadConfig()
	rr := httptest.NewRecorder()
	srv, err := sgbouncewizard.NewServer(cfg)
	handler := http.HandlerFunc(srv.GetAllRulesRoute)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
