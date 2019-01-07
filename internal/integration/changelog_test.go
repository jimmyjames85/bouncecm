package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/sgbouncewizard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ChangelogSuite struct {
	suite.Suite
	srv     *sgbouncewizard.Server
	rr      *httptest.ResponseRecorder
	handler http.HandlerFunc
}

func (suite *ChangelogSuite) SetupTest() {
	dbc := db.Client{Conn: Database}
	suite.srv = &sgbouncewizard.Server{DBClient: &dbc}
	suite.rr = httptest.NewRecorder()
}

func (suite *ChangelogSuite) TestGetAllChangelogsHandler() {
	suite.handler = http.HandlerFunc(suite.srv.GetChangelog) // Change to GetAllChangelogEntries after changelog-CR is merged
	req, err := http.NewRequest("GET", "/changelogs", nil)

	if err != nil {
		suite.T().Errorf("Error in forming request")
	}

	suite.handler.ServeHTTP(suite.rr, req)
	if status := suite.rr.Code; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	assert.Equal(suite.T(), http.StatusOK, suite.rr.Code)
}

// Uncomment after changelog-CR is merged
//
// func TestGetChangelogAcceptance(t *testing.T) {
// 	resp, err := http.Get("http://localhost:3000/changelogs/300")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// }

// func TestChangeLogsGetHandler(t *testing.T) {
// 	rr := httptest.NewRecorder()
// 	dbc := db.Client{Conn: Database}
// 	srv := sgbouncewizard.Server{DBClient: &dbc}
// 	handler := http.HandlerFunc(srv.GetChangelog) // Change to GetAllChangelogEntries after changelog-CR is merged
// 	req, err := http.NewRequest("GET", "/changelogs", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}
// }

func TestChangelogSuite(t *testing.T) {
	suite.Run(t, new(ChangelogSuite))
}
