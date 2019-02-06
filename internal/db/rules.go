package db

import (
	"database/sql"

	"github.com/pkg/errors"

	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// ListRules - Function to pull all rules from db
func (c *Client) GetAllRules() ([]models.BounceRule, error) {
	rules := []models.BounceRule{}
	rows, err := c.Conn.Query("SELECT * FROM bounce_rule")

	if err != nil {
		return nil, errors.Wrap(err, "GetAllRules Query")
	}

	defer rows.Close()

	for rows.Next() {
		br := models.BounceRule{}

		var description sql.NullString
		err := rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &description, &br.BounceAction)

		if description.Valid {
			br.Description = description.String
		}

		if err != nil {
			return nil, errors.Wrap(err, "GetAllRules Scanning")
		}
		rules = append(rules, br)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "GetAllRules Row.Err")
	}

	return rules, nil
}

func (c *Client) GetSingleRule(id int) (*models.BounceRule, error) {
	var br models.BounceRule
	var description sql.NullString
	err := c.Conn.QueryRow("SELECT * From bounce_rule WHERE id = ?", id).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &description, &br.BounceAction)
	if description.Valid {
		br.Description = description.String
	}
	if err != nil {
		return nil, errors.Wrap(err, "GetSingleRule")
	}
	return &br, nil
}

func (c *Client) CreateRule(rule *models.BounceRule) (int, error) {
	id, err := c.Conn.Exec("INSERT INTO bounce_rule(response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?)", rule.ResponseCode, rule.EnhancedCode, rule.Regex, rule.Priority, rule.Description, rule.BounceAction)
	if err != nil {
		return 0, errors.Wrap(err, "CreateRule")
	}

	value , err := id.LastInsertId()

	if err != nil {
		return 0, errors.Wrap(err, "CreateRule LastInserID")
	}

	return int(value), nil
}

func (c *Client) UpdateRule(newRule *models.BounceRule) error {
	_, err := c.Conn.Exec("UPDATE bounce_rule SET id=? , response_code= ? , enhanced_code= ? , regex= ?, priority= ? , description= ?, bounce_action= ? WHERE id= ?", newRule.ID, newRule.ResponseCode, newRule.EnhancedCode, newRule.Regex, newRule.Priority, newRule.Description, newRule.BounceAction, newRule.ID)
	if err != nil {
		return errors.Wrap(err, "UpdateRule")
	}
	return nil
}

func (c *Client) DeleteRule(id int) error {
	_, err := c.Conn.Exec("DELETE FROM bounce_rule WHERE id= ?", id)
	if err != nil {
		return errors.Wrap(err, "DeleteRule")
	}
	return nil
}
