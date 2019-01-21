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
	assert.NoError(suite.T(), err, "Failed to set up change_log table for testing")

	mysql.RegisterLocalFile("testdata/changelog_test.csv")

	res, err := Database.Exec("LOAD DATA LOCAL INFILE '" + "testdata/changelog_test.csv" + "' INTO TABLE changelog FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n'")
	assert.NoError(suite.T(), err, "Failed to get data from file")
	_, err = res.RowsAffected()
	assert.NoError(suite.T(), err, "Failed to get number of rows affected")
}

func (suite *ChangelogSuite) TestGetAllChangelogsHandler() {
	req, err := http.NewRequest("GET", "http://localhost:4000/change_logs", nil)
	assert.NoError(suite.T(), err, "Failed to form GET request")

	res, err := suite.client.Do(req)
	assert.NoError(suite.T(), err, "Failed to send GET request")
	if status := res.StatusCode; status != http.StatusOK {
		suite.T().Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
}

func (suite *ChangelogSuite) TestGetSingleChangelogHandler() {
	resp, err := http.Get("http://localhost:4000/change_logs/400")
	assert.NoError(suite.T(), err, "Failed to send GET request")

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)

	resp, err = http.Get("http://localhost:4000/change_logs/204")
	assert.NoError(suite.T(), err, "Failed to send GET request")
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	changelogs := make([]models.ChangelogEntry, 0)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&changelogs)
	assert.NoError(suite.T(), err, "Failed to marshal struct into JSON")

	assert.Equal(suite.T(), 3, len(changelogs))
	assert.Equal(suite.T(), 204, changelogs[0].ID)
	assert.Equal(suite.T(), "reverted response code and updated enhanced", changelogs[0].Comment)
	assert.Equal(suite.T(), "updated response code", changelogs[1].Comment)
	assert.Equal(suite.T(), "new regex", changelogs[2].Comment)
}

func (suite *ChangelogSuite) TestPostChangelogRoute() {
	resp, err := http.Get("http://localhost:4000/change_logs/604")
	assert.NoError(suite.T(), err, "Failed to send GET request")

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)

	reqBody := models.ChangelogEntry{
		BounceRule: models.BounceRule{
			ID:           204,
			ResponseCode: 403,
			EnhancedCode: "5265126",
			Regex:        "abc*",
			Priority:     0,
			Description:  "RFC5321 Service not available",
			BounceAction: "TRY IT AGAIN",
		},
		UserID:  2,
		Comment: "Fixed the response code (hopefully)",
	}

	preSend, err := json.Marshal(reqBody)
	assert.NoError(suite.T(), err, "Failed to marshal struct into JSON")

	resp, err = http.Post("http://localhost:4000/change_logs", "application/json", bytes.NewBuffer(preSend))
	assert.NoError(suite.T(), err, "Failed to send POST request")

	assert.NotNil(suite.T(), resp)
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	resp, err = http.Get("http://localhost:4000/change_logs/204")
	assert.NoError(suite.T(), err, "Failed to send GET request")

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	changelogs := make([]models.ChangelogEntry, 0)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&changelogs)
	assert.NoError(suite.T(), err, "Failed to marshal struct into JSON")

	assert.Equal(suite.T(), 4, len(changelogs))
	assert.Equal(suite.T(), 204, changelogs[0].ID)
	assert.Equal(suite.T(), "Fixed the response code (hopefully)", changelogs[0].Comment)
	assert.Equal(suite.T(), 520, changelogs[1].ResponseCode)
	assert.Equal(suite.T(), 403, changelogs[0].ResponseCode)
}

func (suite *ChangelogSuite) TearDownTest() {
	_, err := Database.Exec(`TRUNCATE TABLE changelog`)
	assert.NoError(suite.T(), err, "Failed to tear down test data")
}

func (suite *ChangelogSuite) TearDownSuite() {
	_, err := Database.Exec(`DROP TABLE IF EXISTS changelog`)
	assert.NoError(suite.T(), err, "Failed to tear down suite")
}

func TestChangelogSuite(t *testing.T) {
	suite.Run(t, new(ChangelogSuite))
}
