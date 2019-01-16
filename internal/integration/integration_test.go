package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

type DatabaseTestSuite struct {
	suite.Suite
	client *http.Client
}

func (suite *DatabaseTestSuite) SetupSuite() {
	suite.client = &http.Client{}
}

func (suite *DatabaseTestSuite) SetupTest() {
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

	if err != nil {
		suite.T().Fatalf("Failed to setup for test\nError: %v", err)
	}

	mysql.RegisterLocalFile("testdata/bounce_rules.csv")

	res, err := Database.Exec("LOAD DATA LOCAL INFILE '" + "testdata/bounce_rules.csv" + "' INTO TABLE bounce_rule FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n'")
	if err != nil {
		suite.T().Fatalf("Failed to load data from file\nError: %v", err)
	}
	inserted, err := res.RowsAffected()
	if err != nil {
		suite.T().Fatalf("Failed to get inserted rows\nError: %v", err)
	}
	if inserted != 295 {
		suite.T().Fatalf("Expected %d rows in bounce_rule table, got %d", 295, inserted)
	}

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
	if err != nil {
		suite.T().Fatalf("Failed to setup for test\nError: %v", err)
	}

	mysql.RegisterLocalFile("testdata/changelog_test.csv")

	res, err = Database.Exec("LOAD DATA LOCAL INFILE '" + "testdata/changelog_test.csv" + "' INTO TABLE changelog FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n'")
	if err != nil {
		suite.T().Fatalf("Failed to load data from file\nError: %v", err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		suite.T().Fatalf("Failed to get inserted rows\nError: %v", err)
	}
}

func (suite *DatabaseTestSuite) TestBounceRuleDatabaseSelectSingle() {
	var br models.BounceRule
	err := Database.QueryRow(`SELECT * From bounce_rule WHERE id = 174`).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)

	if err != nil {
		suite.T().Errorf("Database query failed")
	}

	assert.Equal(suite.T(), 174, br.ID)
	assert.Equal(suite.T(), 450, br.ResponseCode)
	assert.Equal(suite.T(), "", br.EnhancedCode)
	assert.Equal(suite.T(), "", br.Regex)
	assert.Equal(suite.T(), 0, br.Priority)
	assert.Equal(suite.T(), "RFC5321 Mailbox unavailable", br.Description)
	assert.Equal(suite.T(), "retry", br.BounceAction)
}

func (suite *DatabaseTestSuite) TestBounceRuleDatabaseSelectMultiple() {
	var br models.BounceRule
	rows, err := Database.Query(`SELECT * From bounce_rule WHERE response_code = 550`)

	if err != nil {
		suite.T().Errorf("Database query failed")
	}

	var rowCount = 0
	for rows.Next() {
		if err := rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction); err != nil {
			suite.T().Fatalf("Failed to read row into struct\nError: %v", err)
		}
		assert.Equal(suite.T(), 550, br.ResponseCode)
		rowCount++
	}
	assert.Equal(suite.T(), 81, rowCount)
}

func (suite *DatabaseTestSuite) TestBounceRuleDatabaseInsert() {
	var br models.BounceRule
	err := Database.QueryRow(`SELECT * From bounce_rule WHERE id = 507`).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)

	if err == nil {
		suite.T().Errorf("Database query failed")
	}

	assert.Equal(suite.T(), 0, br.ID)
	assert.Equal(suite.T(), "", br.BounceAction)

	res, err := Database.Exec("INSERT INTO bounce_rule(response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?)", 504, "123", "445", 0, "RF()", "try again")

	if err != nil {
		suite.T().Errorf("Failed to execute INSERT query %v\nError: ", err)
	}

	insertedID, err := res.LastInsertId()

	if err != nil {
		suite.T().Errorf("Failed to get last inserted ID from query\nError: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		suite.T().Errorf("Failed to get rows affected by query\nError: %v", err)
	}

	assert.Equal(suite.T(), int64(507), insertedID)
	assert.Equal(suite.T(), int64(1), rowsAffected)

	err = Database.QueryRow(`SELECT * From bounce_rule WHERE id = 507`).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
	assert.Equal(suite.T(), 507, br.ID)
	assert.Equal(suite.T(), 504, br.ResponseCode)
	assert.Equal(suite.T(), "123", br.EnhancedCode)
	assert.Equal(suite.T(), "445", br.Regex)
	assert.Equal(suite.T(), 0, br.Priority)
	assert.Equal(suite.T(), "RF()", br.Description)
	assert.Equal(suite.T(), "try again", br.BounceAction)
}

func (suite *DatabaseTestSuite) TestBounceRuleDatabaseDelete() {
	res, err := Database.Exec(`DELETE FROM bounce_rule where id = 507`)

	if err != nil {
		suite.T().Fatalf("Erroneous test data detected\nError: %v", err)
	}

	rowsDeleted, err := res.RowsAffected()

	if err != nil {
		suite.T().Fatalf("Failed to retrieve number of rows affected by DELETE\nError: %v", err)
	}

	if rowsDeleted != 0 {
		suite.T().Fatalf("Errorneous test data detected\nError: %v", err)
	}

	_, err = Database.Exec(`DELETE FROM bounce_rule where id = 507`)

	if err != nil {
		suite.T().Fatalf("Erroneous test data detected\nError: %v", err)
	}
}

func (suite *DatabaseTestSuite) TearDownTest() {
	_, err := Database.Exec(`TRUNCATE TABLE bounce_rule`)
	if err != nil {
		suite.T().Fatalf("Failed to tear down test data\nError: %v", err)
	}
	_, err = Database.Exec(`TRUNCATE TABLE changelog`)
	if err != nil {
		suite.T().Fatalf("Failed to tear down test data\nError: %v", err)
	}
}

func (suite *DatabaseTestSuite) TearDownSuite() {
	_, err := Database.Exec(`DROP TABLE IF EXISTS bounce_rule`)
	if err != nil {
		suite.T().Fatalf("Failed to tear down test data\nError: %v", err)
	}
	_, err = Database.Exec(`DROP TABLE IF EXISTS changelog`)
	if err != nil {
		suite.T().Fatalf("Failed to tear down test data\nError: %v", err)
	}
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
