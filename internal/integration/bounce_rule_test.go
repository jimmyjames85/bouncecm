package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-sql-driver/mysql"

	"github.com/jimmyjames85/bouncecm/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BounceRuleSuite struct {
	suite.Suite

	handler http.HandlerFunc
	client  *http.Client
}

func (suite *BounceRuleSuite) SetupSuite() {
	suite.client = &http.Client{}
}

func (suite *BounceRuleSuite) SetupTest() {
	res, err := Database.Exec(`
		CREATE TABLE IF NOT EXISTS bounce_rule (
			id smallint(5) unsigned NOT NULL AUTO_INCREMENT,
			response_code smallint(5) unsigned NOT NULL DEFAULT '0',
			enhanced_code varchar(16) NOT NULL DEFAULT '',
			regex varchar(255) NOT NULL DEFAULT '',
			priority tinyint(3) unsigned NOT NULL DEFAULT '0',
			description varchar(255) DEFAULT NULL,
			bounce_action varchar(255) NOT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY bounce_rule_components (response_code,enhanced_code,regex)
		) ENGINE=InnoDB DEFAULT CHARSET=latin1;
	`)

	if err != nil {
		suite.T().Fatalf("Failed to setup for test\nError: %v", err)
	}

	mysql.RegisterLocalFile("../../bounce_rule_test.txt")

	res, err = Database.Exec("LOAD DATA LOCAL INFILE '" + "../../bounce_rule_test.txt" + "' INTO TABLE bounce_rule FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n'")
	if err != nil {
		suite.T().Fatalf("Failed to load data from file\nError: %v", err)
	}
	inserted, err := res.RowsAffected()
	if err != nil {
		suite.T().Fatalf("Failed to get inserted rows\nError: %v", err)
	}
	if inserted != 49 {
		suite.T().Fatalf("Expected %d rows in bounce_rule table, got %d", 49, inserted)
	}
}

func (suite *BounceRuleSuite) TestGetAllBounceRulesHandler() {
	req, err := http.NewRequest("GET", "http://localhost:4000/bounce_rules", nil)
	if err != nil {
		suite.T().Fatalf("Error in forming request\nError: %v", err)
	}

	res, err := suite.client.Do(req)
	if err != nil {
		suite.T().Errorf("Failed to execute request\nError: %v", err)
	}
	if status := res.StatusCode; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
}

func (suite *BounceRuleSuite) TestGetSingleBounceRuleHandler() {
	resp, err := http.Get("http://localhost:4000/bounce_rules/900")
	if err != nil {
		suite.T().Errorf("Request failed")
	}

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)

	resp, err = http.Get("http://localhost:4000/bounce_rules/180")
	if err != nil {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	// assert contents
}

func (suite *BounceRuleSuite) TestPostBounceRuleHandler() {
	resp, err := http.Get("http://localhost:4000/bounce_rules/506")
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

	resp, err = http.Post("http://localhost:4000/bounce_rules", "application/json", bytes.NewBuffer(preSend))
	if err != nil {
		suite.T().Errorf("POST to route failed")
	}

	assert.NotNil(suite.T(), resp)
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	// Redo GET request to confirm rule is in database
	// Delete Rule
}

func (suite *BounceRuleSuite) TestDeleteBounceRuleHandler() {
	req, err := http.NewRequest("DELETE", "http://localhost:4000/bounce_rules/900", nil)
	if err != nil {
		suite.T().Errorf("Creation of request failed")
	}
	res, err := suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusNotFound {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound) // switch to http.StatusNotFound when PR merged
	}

	req, err = http.NewRequest("DELETE", "http://localhost:4000/bounce_rules/200", nil)
	if err != nil {
		suite.T().Errorf("Failed to create DELETE request\nError: %v", err)
	}

	res, err = suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	req, err = http.NewRequest("DELETE", "http://localhost:4000/bounce_rules/200", nil)
	if err != nil {
		suite.T().Errorf("Failed to create DELETE request\nError: %v", err)
	}

	res, err = suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusNotFound {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

}

func (suite *BounceRuleSuite) TestUpdateBounceRuleHandler() {
	reqBody := models.BounceRule{
		ID: 800, ResponseCode: 403, EnhancedCode: "111", Regex: "asfba", Priority: 0, Description: "Test Update", BounceAction: "Do Nothing",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "http://localhost:4000/bounce_rules/800", bytes.NewBuffer(jsonBody))
	if err != nil {
		suite.T().Errorf("Creation of request failed")
	}

	res, err := suite.client.Do(req)
	if err != nil {
		suite.T().Errorf("Request had failed\nError: %v", err)
	}
	if status := res.StatusCode; status != http.StatusNotFound {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	req, err = http.NewRequest("PUT", "http://localhost:4000/bounce_rules/180", bytes.NewBuffer(jsonBody))
	if err != nil {
		suite.T().Errorf("Failed to create UPDATE request\nError: %v", err)
	}

	res, err = suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func (suite *BounceRuleSuite) TearDownTest() {
	_, err := Database.Exec(`DELETE FROM bounce_rule`)
	if err != nil {
		suite.T().Fatalf("Failed to tear down test data\nError: %v", err)
	}
}

func TestBounceRuleSuite(t *testing.T) {
	suite.Run(t, new(BounceRuleSuite))
}
