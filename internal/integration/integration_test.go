package integration

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

func TestGetAllBounceRules(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/bounce_rules")
	log.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestGetBounceRule(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/bounce_rules/173")
	assert.Nil(t, err)
	assert.NotNil(t, resp)

}

func TestPostBounceRule(t *testing.T) {
	reqBody := map[string]interface{}{
		"ResponseCode": 421,
		"EnhancedCode": "",
		"Regex":        "",
		"Priority":     0,
		"Description":  "RFC5321 Service not available",
		"BounceAction": "retry",
	}
	preSend, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:3000/bounce_rules", "application/json", bytes.NewBuffer(preSend))
	if err != nil {
		log.Fatalln(err)
	}
	assert.NotNil(t, resp)
}

// func TestDeleteBounceRule(t *testing.T) {
// 	cfg, err := config.LoadConfig()
// 	rr := httptest.NewRecorder()
// 	srv, err := sgbouncewizard.NewServer(cfg)
// 	handler := http.HandlerFunc(srv.DeleteRuleRoute)

// 	req, err := http.NewRequest("DELETE", "/bounce_rules/600", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusNotFound {
// 		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
// 	}
// 	expected := "404 page not found\n"
// 	assert.Equal(t, expected, rr.Body.String(), "Response body differs")

// 	req, err = http.NewRequest("DELETE", "/bounce_rules/504", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	req, err = http.NewRequest("DELETE", "/bounce_rules/504", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if status := rr.Code; status != http.StatusNotFound {
// 		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// }

func TestGetAllChangelogs(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/changelogs")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

// Uncomment after changelog-CR is merged
//
// func TestGetChangelog(t *testing.T) {
// 	resp, err := http.Get("http://localhost:3000/changelogs/300")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// }

func TestChangeLogsGetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/changelogs", nil)

	if err != nil {
		t.Fatal(err)
	}
	cfg, err := config.LoadConfig()
	rr := httptest.NewRecorder()
	srv, err := sgbouncewizard.NewServer(cfg)
	handler := http.HandlerFunc(srv.GetChangelog) // Change to GetAllChangelogEntries after changelog-CR is merged
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestBounceRuleGetAllHandler(t *testing.T) {
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

func TestBounceRuleGetSingleHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/bounce_rules/180", nil)

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
