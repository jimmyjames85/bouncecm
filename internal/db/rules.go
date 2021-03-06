package db

import (
	"database/sql"
	"github.com/pkg/errors"
	"strconv"

	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// ListRules - Function to pull all rules from db
func (c *Client) GetAllRulesLimited(offset int, limit int) ([]models.BounceRule, error) {
	rules := []models.BounceRule{}
	rows, err := c.Conn.Query("SELECT * FROM bounce_rule LIMIT ?,?", offset, limit)

	if err != nil {
		return nil, errors.Wrap(err, "GetAllRulesLimited Query")
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
			return nil, errors.Wrap(err, "GetAllRulesLimited Scanning")
		}
		rules = append(rules, br)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "GetAllRulesLimited Row.Err")
	}

	return rules, nil
}

func (c *Client) GetAllRulesFiltered(offset int, limit int, filterby string, option string) ([]models.BounceRule, error) {
	rules := []models.BounceRule{}
	var rows *sql.Rows
	var err error
	if filterby == "bounce_action" {
		rows, err = c.Conn.Query("SELECT * FROM bounce_rule where bounce_action = ?  LIMIT ?,?", option, offset, limit)
	}

	if filterby == "priority" {
		priority, err := strconv.Atoi( option)
		if err != nil {
			return nil, errors.Wrap(err, "GetAllRulesFiltered Priority needs to be an int")
		}

		rows, err = c.Conn.Query("SELECT * FROM bounce_rule where priority = ?  LIMIT ?,?",priority, offset, limit)
	}

	
	if filterby == "response_code" {
		response_code, err := strconv.Atoi(option)
		if err != nil {
			return nil, errors.Wrap(err, "GetAllRulesFiltered response_code needs to be an int")
		}

		rows, err = c.Conn.Query("SELECT * FROM bounce_rule where response_code = ?  LIMIT ?,?",response_code, offset, limit)
	}

		
	if filterby == "description" {
		description := "%" + option + "%"
		rows, err = c.Conn.Query("SELECT * FROM bounce_rule where description LIKE ?  LIMIT ?,?",description, offset, limit)
	}
	
	if err != nil {
		return nil, errors.Wrap(err, "GetAllRulesFiltered Query")
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
			return nil, errors.Wrap(err, "GetAllRulesFiltered Scanning")
		}
		rules = append(rules, br)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "GetAllRulesFiltered Row.Err")
	}

	return rules, nil
}


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
