package integration

import (
	"net/http"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ChangelogSuite struct {
	suite.Suite
	client *http.Client
}

func (suite *ChangelogSuite) SetupSuite() {
	suite.client = &http.Client{}
}

func (suite *ChangelogSuite) SetupTest() {
	_, err := Database.Exec(`
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
	if err != nil {
		suite.T().Fatalf("Failed to setup for test\nError: %v", err)
	}

	mysql.RegisterLocalFile("testdata/changelog_test.csv")

	res, err := Database.Exec("LOAD DATA LOCAL INFILE '" + "testdata/changelog_test.csv" + "' INTO TABLE changelog FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n'")
	if err != nil {
		suite.T().Fatalf("Failed to load data from file\nError: %v", err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		suite.T().Fatalf("Failed to get inserted rows\nError: %v", err)
	}
}

func (suite *ChangelogSuite) TestGetAllChangelogsHandler() {
	req, err := http.NewRequest("GET", "http://localhost:4000/change_logs", nil)
	if err != nil {
		suite.T().Errorf("Error in forming request")
	}

	res, err := suite.client.Do(req)
	if err != nil {
		suite.T().Errorf("GET request was not able to be completed\nError: %v", err)
	}
	if status := res.StatusCode; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
}

// Uncomment after changelog-CR is merged
//
// func (suite *ChangelogSuite) TestGetSingleChangelogHandler() {
// 	resp, err := http.Get("http://localhost:3000/change_logs/300")
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

// func (suite *ChangelogSuite) TestPostChangelogRoute() {
// 	resp, err := http.Get("http://localhost:4000/change_logs/604")
// 	if err != nil {
// 		suite.T().Errorf("GET requested failed")
// 	}

// 	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)

// 	reqBody := map[string]interface{}{
// 		"lastId": 12,
// 		"UserID": 2,
// 		"Comment": "Fixed the response code (hopefully)"
// 		"ResponseCode": 403,
// 		"EnhancedCode": "5265126",
// 		"Regex":        "1212121",
// 		"Priority":     0,
// 		"Description":  "RFC5321 Service not available",
// 		"BounceAction": "TRY IT AGAIN",
// 	}
// 	preSend, err := json.Marshal(reqBody)
// 	if err != nil {
// 		suite.T().Errorf("Formatting of JSON incorrect")
// 	}

// 	resp, err = http.Post("http://localhost:4000/change_logs", "application/json", bytes.NewBuffer(preSend))
// 	if err != nil {
// 		suite.T().Errorf("POST to route failed")
// 	}

// 	assert.NotNil(suite.T(), resp)
// 	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

// 	resp, err = http.Get("http://localhost:4000/change_logs/507")
// 	if err != nil {
// 		suite.T().Errorf("GET requested failed")
// 	}

// 	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
// }

func (suite *ChangelogSuite) TearDownTest() {
	_, err := Database.Exec(`TRUNCATE TABLE changelog`)
	if err != nil {
		suite.T().Fatalf("Failed to tear down test data\nError: %v", err)
	}
}

func (suite *ChangelogSuite) TearDownSuite() {
	_, err := Database.Exec(`DROP TABLE IF EXISTS changelog`)
	if err != nil {
		suite.T().Fatalf("Failed to remove test table from database\nError: %v", err)
	}
}

func TestChangelogSuite(t *testing.T) {
	suite.Run(t, new(ChangelogSuite))
}
