package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/sgbouncewizard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BounceRuleSuite struct {
	suite.Suite
	srv     *sgbouncewizard.Server
	rr      *httptest.ResponseRecorder
	handler http.HandlerFunc
}

func (suite *BounceRuleSuite) SetupTest() {
	dbc := db.Client{Conn: Database}
	suite.srv = &sgbouncewizard.Server{DBClient: &dbc}
	suite.rr = httptest.NewRecorder()

}

func (suite *BounceRuleSuite) TestGetAllBounceRulesHandler() {
	suite.handler = http.HandlerFunc(suite.srv.GetAllRulesRoute)
	req, err := http.NewRequest("GET", "/bounce_rules", nil)

	if err != nil {
		suite.T().Errorf("Error in forming request")
	}

	suite.handler.ServeHTTP(suite.rr, req)
	if status := suite.rr.Code; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	assert.Equal(suite.T(), http.StatusOK, suite.rr.Code)
}

func (suite *BounceRuleSuite) TestGetSingleBounceRuleHandler() {
	resp, err := http.Get("http://localhost:3000/bounce_rules/900")

	if err != nil {
		suite.T().Errorf("Request failed")
	}

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)

	resp, err = http.Get("http://localhost:3000/bounce_rules/180")

	if err != nil {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	// assert contents
}

func (suite *BounceRuleSuite) TestPostBounceRuleHandler() {
	resp, err := http.Get("http://localhost:3000/bounce_rules/506")

	if err != nil {
		suite.T().Errorf("GET requested failed")
	}

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)

	reqBody := map[string]interface{}{
		"ResponseCode": 421,
		"EnhancedCode": "5235123",
		"Regex":        "1111111",
		"Priority":     0,
		"Description":  "RFC5321 Service not available",
		"BounceAction": "retry",
	}
	preSend, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Errorf("Formatting of JSON incorrect")
	}

	resp, err = http.Post("http://localhost:3000/bounce_rules", "application/json", bytes.NewBuffer(preSend))
	if err != nil {
		suite.T().Errorf("POST to route failed")
	}

	assert.NotNil(suite.T(), resp)
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	// Redo GET request to confirm rule is in database
	// Delete Rule
}

// func (suite *BounceRuleSuite) TestDeleteBounceRuleHandler() {
// 	req, err := http.NewRequest("DELETE", "/bounce_rules/900", nil)

// 	if err != nil {
// 		suite.T().Errorf("Creation of request failed")
// 	}

// 	suite.handler.ServeHTTP(suite.rr, req)

// 	if status := suite.rr.Code; status != http.StatusNotFound {
// 		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
// 	}
// 	assert.Equal(suite.T(), http.StatusNotFound, suite.rr.Code)

// }

// func (suite *BounceRuleSuite) TestUpdateBounceRuleHandler() {
// 	req, err := http.NewRequest("UPDATE", "/bounce_rules/900", nil)

// 	if err != nil {
// 		suite.T().Errorf("Creation of request failed")
// 	}

// 	suite.handler.ServeHTTP(suite.rr, req)

// 	if status := suite.rr.Code; status != http.StatusNotFound {
// 		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
// 	}
// 	assert.Equal(suite.T(), http.StatusNotFound, suite.rr.Code)
// }

// Commented until changelog-CR PR merged
//
// func TestDeleteBounceRuleAcceptance(t *testing.T) {
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

func TestBounceRuleSuite(t *testing.T) {
	suite.Run(t, new(BounceRuleSuite))
}
