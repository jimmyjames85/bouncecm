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
	_, err := Database.Exec(`
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

	assert.NoError(suite.T(), err, "Failed to setup table")

	// Required, as POSTing to a bounce_rule route also creates a changelog
	_, err = Database.Exec(`
		CREATE TABLE IF NOT EXISTS changelog (
			rule_id smallint(5) unsigned NOT NULL,
			user_id smallint(5) unsigned NOT NULL,
			comment varchar(255) NOT NULL,
			created_at int(11) NOT NULL,
			response_code smallint(5) unsigned NOT NULL DEFAULT '0',
			enhanced_code varchar(16) NOT NULL DEFAULT '',
			regex varchar(255) NOT NULL DEFAULT '',
			priority tinyint(3) unsigned NOT NULL DEFAULT '0',
			description varchar(255) DEFAULT NULL,
			bounce_action varchar(255) NOT NULL,
			PRIMARY KEY (created_at)
	  	) ENGINE=InnoDB DEFAULT CHARSET=latin1;`)
	assert.NoError(suite.T(), err, "Failed to set up change_log table for testing")

	mysql.RegisterLocalFile("testdata/bounce_rules.csv")

	res, err := Database.Exec("LOAD DATA LOCAL INFILE '" + "testdata/bounce_rules.csv" + "' INTO TABLE bounce_rule FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n'")
	assert.NoError(suite.T(), err, "Failed to load data from the table")

	inserted, err := res.RowsAffected()
	assert.NoError(suite.T(), err, "Failed to get rows from table")
	if inserted <= 0 {
		suite.T().Fatalf("Expected rows in bounce_rule table, got 0\n Error: %v", err)
	}
}

func (suite *BounceRuleSuite) TestGetAllBounceRulesHandler() {
	req, err := http.NewRequest("GET", "http://localhost:4000/bounce_rules", nil)
	assert.NoError(suite.T(), err, "Failed to form request")

	res, err := suite.client.Do(req)
	assert.NoError(suite.T(), err, "Failed in executing request")
	if status := res.StatusCode; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
}

func (suite *BounceRuleSuite) TestGetSingleBounceRuleHandler() {
	resp, err := http.Get("http://localhost:4000/bounce_rules/900")
	assert.NoError(suite.T(), err, "Failed to form request")

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)

	resp, err = http.Get("http://localhost:4000/bounce_rules/180")
	assert.NoError(suite.T(), err, "Failed to form request")

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	// assert contents

	var want models.BounceRule
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&want)
	assert.NoError(suite.T(), err, "Failed to marshal struct into JSON")

	assert.Equal(suite.T(), 180, want.ID)
	assert.Equal(suite.T(), 501, want.ResponseCode)
}

func (suite *BounceRuleSuite) TestPostBounceRuleHandler() {
	resp, err := http.Get("http://localhost:4000/bounce_rules/507")
	assert.NoError(suite.T(), err, "Failed to form request")

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)

	have := models.BounceRule{
		ID:           507,
		ResponseCode: 421,
		EnhancedCode: "5235123",
		Regex:        "1111132",
		Priority:     0,
		Description:  "RFC5321 Service not available",
		BounceAction: "try it again homie",
	}

	preSend, err := json.Marshal(have)

	assert.NoError(suite.T(), err, "Failed to marshal JSON")

	resp, err = http.Post("http://localhost:4000/bounce_rules", "application/json", bytes.NewBuffer(preSend))
	assert.NoError(suite.T(), err, "Failed to send POST request")

	assert.NotNil(suite.T(), resp)
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	resp, err = http.Get("http://localhost:4000/bounce_rules/507")
	assert.NoError(suite.T(), err, "Failed to send GET request")

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	var want models.BounceRule
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&want)
	assert.NoError(suite.T(), err, "Failed to marshal struct into JSON")

	assert.Equal(suite.T(), have, want)

}

func (suite *BounceRuleSuite) TestDeleteBounceRuleHandler() {
	req, err := http.NewRequest("DELETE", "http://localhost:4000/bounce_rules/900", nil)
	assert.NoError(suite.T(), err, "Failed to form request")
	res, err := suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusNotFound {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound) // switch to http.StatusNotFound when PR merged
	}

	req, err = http.NewRequest("DELETE", "http://localhost:4000/bounce_rules/200", nil)
	assert.NoError(suite.T(), err, "Failed to form DELETE request")

	res, err = suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	req, err = http.NewRequest("DELETE", "http://localhost:4000/bounce_rules/200", nil)
	assert.NoError(suite.T(), err, "Failed to form DELETE request")

	res, err = suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusNotFound {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	req, err = http.NewRequest("GET", "http://localhost:4000/bounce_rules/200", nil)
	assert.NoError(suite.T(), err, "Failed to form GET request")

	res, err = suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusNotFound {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

}

func (suite *BounceRuleSuite) TestUpdateBounceRuleHandler() {
	have := models.BounceRule{
		ID: 180, ResponseCode: 403, EnhancedCode: "111", Regex: "asfba", Priority: 0, Description: "Test Update", BounceAction: "Do Nothing",
	}
	jsonBody, _ := json.Marshal(have)

	req, err := http.NewRequest("PUT", "http://localhost:4000/bounce_rules/800", bytes.NewBuffer(jsonBody))
	assert.NoError(suite.T(), err, "Failed to form PUT request")

	res, err := suite.client.Do(req)
	assert.NoError(suite.T(), err, "Failed to send PUT request")

	if status := res.StatusCode; status != http.StatusNotFound {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	req, err = http.NewRequest("PUT", "http://localhost:4000/bounce_rules/180", bytes.NewBuffer(jsonBody))
	assert.NoError(suite.T(), err, "Failed to form PUT request")

	res, err = suite.client.Do(req)
	if status := res.StatusCode; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	resp, err := http.Get("http://localhost:4000/bounce_rules/180")
	assert.NoError(suite.T(), err, "Failed to send GET request")

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	var want models.BounceRule
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&want)
	assert.NoError(suite.T(), err, "Failed to marshal struct into JSON")

	assert.Equal(suite.T(), have, want)

}

func (suite *BounceRuleSuite) TearDownTest() {
	_, err := Database.Exec(`TRUNCATE TABLE bounce_rule`)
	assert.NoError(suite.T(), err, "Failed to tear down database")
	_, err = Database.Exec(`TRUNCATE TABLE changelog`)
	assert.NoError(suite.T(), err, "Failed to tear down database")
}

func (suite *BounceRuleSuite) TearDownSuite() {
	_, err := Database.Exec(`DROP TABLE IF EXISTS bounce_rule`)
	assert.NoError(suite.T(), err, "Failed to tear down suite")
	_, err = Database.Exec(`DROP TABLE IF EXISTS changelog`)
	assert.NoError(suite.T(), err, "Failed to tear down suite")
}

func TestBounceRuleSuite(t *testing.T) {
	suite.Run(t, new(BounceRuleSuite))
}
