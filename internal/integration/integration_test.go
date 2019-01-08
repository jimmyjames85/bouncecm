package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// Refactor into Suite to use SetupTest and TearDownTest

func TestBounceRuleDatabaseSelect(t *testing.T) {
	var br models.BounceRule
	err := Database.QueryRow(`SELECT * From bounce_rule WHERE id = 174`).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)

	if err != nil {
		t.Errorf("Database query failed")
	}

	assert.Equal(t, 174, br.ID)
	assert.Equal(t, 450, br.ResponseCode)
	assert.Equal(t, "", br.EnhancedCode)
	assert.Equal(t, "", br.Regex)
	assert.Equal(t, 0, br.Priority)
	assert.Equal(t, "RFC5321 Mailbox unavailable", br.Description)
	assert.Equal(t, "retry", br.BounceAction)
}

func TestBounceRuleDatabaseInsert(t *testing.T) {
	var br models.BounceRule
	err := Database.QueryRow(`SELECT * From bounce_rule WHERE id = 505`).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)

	if err == nil {
		t.Errorf("Database query failed")
	}

	assert.Equal(t, 0, br.ID)
	assert.Equal(t, "", br.BounceAction)

	res, err := Database.Exec("INSERT INTO bounce_rule(response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?)", 504, "123", "445", 0, "RF()", "try again")

	if err != nil {
		t.Errorf("Failed to execute INSERT query %v\nError: ", err)
	}

	insertedID, err := res.LastInsertId()

	if err != nil {
		t.Errorf("Failed to get last inserted ID from query\nError: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		t.Errorf("Failed to get rows affected by query\nError: %v", err)
	}

	assert.Equal(t, int64(505), insertedID)
	assert.Equal(t, int64(1), rowsAffected)

	err = Database.QueryRow(`SELECT * From bounce_rule WHERE id = 505`).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
	assert.Equal(t, 505, br.ID)
	assert.Equal(t, 504, br.ResponseCode)
	assert.Equal(t, "123", br.EnhancedCode)
	assert.Equal(t, "445", br.Regex)
	assert.Equal(t, 0, br.Priority)
	assert.Equal(t, "RF()", br.Description)
	assert.Equal(t, "try again", br.BounceAction)

	deleted, err := Database.Exec(`DELETE FROM bounce_rule WHERE id = 505`)

	if err != nil {
		t.Fatalf("Failed to remove test data\nError: %v", err)
	}

	deletedRows, err := deleted.RowsAffected()

	if err != nil {
		t.Errorf("Failed to get number of deleted rows")
	}

	if deletedRows != 1 {
		t.Fatalf("Failed to remove test data\nError: %v", err)
	}
}

func TestBounceRuleDatabaseDelete(t *testing.T) {
	res, err := Database.Exec(`DELETE FROM bounce_rule where id = 505`)

	if err != nil {
		t.Fatalf("Erroneous test data detected\nError: %v", err)
	}

	rowsDeleted, err := res.RowsAffected()

	if err != nil {
		t.Fatalf("Failed to retrieve number of rows affected by DELETE\nError: %v", err)
	}

	if rowsDeleted != 0 {
		t.Fatalf("Errorneous test data detected\nError: %v", err)
	}

	_, err = Database.Exec(`DELETE FROM bounce_rule where id = 506`)

	if err != nil {
		t.Fatalf("Erroneous test data detected\nError: %v", err)
	}

	_, err = Database.Exec("INSERT INTO bounce_rule(response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?)", 505, "124", "testing", 0, "RFC1", "try again")

	if err != nil {
		t.Fatalf("Failed to insert into table\nError: %v", err)
	}

	_, err = Database.Exec(`DELETE FROM bounce_rule where id = 506`)

	if err != nil {
		t.Fatalf("Failed to delete from table")
	}
}
