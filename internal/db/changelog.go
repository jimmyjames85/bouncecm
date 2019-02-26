package db

import (
	"database/sql"
	"regexp"
	"time"

	"github.com/pkg/errors"

	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// Changelog - Function to pull all rules from db
func (c *Client) GetAllChangelogEntriesLimited( offset int, limit int) ([]models.ChangelogEntry, error) {
	rules := []models.ChangelogEntry{}

	rows, err := c.Conn.Query("SELECT * FROM changelog order by created_at Desc LIMIT ?,?", offset, limit)

	if err != nil {
		return nil, errors.Wrap(err, "Changelog with limit and offset Query Error")
	}

	defer rows.Close()

	for rows.Next() {
		br := models.ChangelogEntry{}
		var description sql.NullString
		err = rows.Scan(&br.ID, &br.UserID, &br.Comment, &br.CreatedAt, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &description, &br.BounceAction , &br.Operation)


		if description.Valid {
			br.Description = description.String
		}

		if err != nil {
			return nil, errors.Wrap(err, "Changelog with limit and offset Row Scan")
		}

		rules = append(rules, br)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "Changelog with limit and offset Row Error")
	}

	return rules, nil
}

func (c *Client) GetAllChangelogEntries() ([]models.ChangelogEntry, error) {
	rules := []models.ChangelogEntry{}

	rows, err := c.Conn.Query("SELECT * FROM changelog")

	if err != nil {
		return nil, errors.Wrap(err, "Changelog Query Error")
	}

	defer rows.Close()

	for rows.Next() {
		br := models.ChangelogEntry{}
		var description sql.NullString
		err = rows.Scan(&br.ChangelogID, &br.ID, &br.UserID, &br.Comment, &br.CreatedAt, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &description, &br.BounceAction, &br.Operation)

		if description.Valid {
			br.Description = description.String
		}

		if err != nil {
			return nil, errors.Wrap(err, "Changelog Row Scan")
		}

		rules = append(rules, br)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "Changelog Row Error")
	}

	return rules, nil
}

func (c *Client) GetChangeLogByIdLimited(id int, offset int, limit int) ([]models.ChangelogEntry, error) {

	var rows *sql.Rows
	var err error

	rows, err = c.Conn.Query("SELECT * From changelog WHERE rule_id = ?  ORDER BY created_at DESC LIMIT ?,?", id, offset, limit)


	if err != nil {
		return nil, errors.Wrap(err, "GetChangeLogById Query")
	}

	rules := []models.ChangelogEntry{}

	defer rows.Close()

	for rows.Next() {
		cl := models.ChangelogEntry{}
		var description sql.NullString
		err = rows.Scan(&cl.ID, &cl.UserID, &cl.Comment, &cl.CreatedAt, &cl.ResponseCode, &cl.EnhancedCode, &cl.Regex, &cl.Priority, &description, &cl.BounceAction, &cl.Operation)
		if err != nil {
			return nil, errors.Wrap(err, "GetChangeLogById Row Scan")
		}
		if description.Valid {
			cl.Description = description.String
		}
		rules = append(rules, cl)
	}

	if len(rules) <= 0 {
		emptyChangelogError := errors.New("sql: no rows in result set")
		return nil, errors.Wrap(emptyChangelogError, "GetChangeLogById")
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "GetAllRules Row.Err")
	}

	return rules, nil
}


func (c *Client) GetChangeLogById(id int) ([]models.ChangelogEntry, error) {

	var rows *sql.Rows
	var err error

	rows, err = c.Conn.Query("SELECT * From changelog WHERE rule_id = ?  ORDER BY created_at DESC ", id)


	if err != nil {
		return nil, errors.Wrap(err, "GetChangeLogById Query")
	}

	rules := []models.ChangelogEntry{}

	defer rows.Close()

	for rows.Next() {
		cl := models.ChangelogEntry{}
		var description sql.NullString
		err = rows.Scan(&cl.ChangelogID, &cl.ID, &cl.UserID, &cl.Comment, &cl.CreatedAt, &cl.ResponseCode, &cl.EnhancedCode, &cl.Regex, &cl.Priority, &description, &cl.BounceAction, &cl.Operation)
		if err != nil {
			return nil, errors.Wrap(err, "GetChangeLogById Row Scan")
		}
		if description.Valid {
			cl.Description = description.String
		}
		rules = append(rules, cl)
	}

	if len(rules) <= 0 {
		emptyChangelogError := errors.New("sql: no rows in result set")
		return nil, errors.Wrap(emptyChangelogError, "GetChangeLogById")
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "GetAllRules Row.Err")
	}

	return rules, nil
}

func (c *Client) CreateChangeLogEntry(lastId int, entry *models.ChangelogEntry) error {
	_, err := regexp.Compile(entry.Regex)

	if err != nil {
		return errors.Wrap(err, "Invalid Regex")
	}
	_, err = c.Conn.Exec("INSERT INTO changelog(rule_id,user_id,comment,created_at,response_code,enhanced_code,regex,priority,description,bounce_action,operation) VALUES(?,?,?,?,?,?,?,?,?,?,?)", lastId, entry.UserID, entry.Comment, int32(time.Now().Unix()), entry.ResponseCode, entry.EnhancedCode, entry.Regex, entry.Priority, entry.Description, entry.BounceAction, entry.Operation)
	if err != nil {
		return errors.Wrap(err, "CreateChangeLogEntry")
	}
	return nil
}
