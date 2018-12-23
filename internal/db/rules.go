package db

import (
	"database/sql"
	"log"
	"github.com/pkg/errors"
	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// ListRules - Function to pull all rules from db
func (c *Client) GetAllRules() (*models.RulesObject, error) {
	rules := []models.BounceRule{}
	rows, err := c.Conn.Query("SELECT * FROM bounce_rule")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		br := models.BounceRule{}

		var description sql.NullString
		err := rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &description, &br.BounceAction)

		if description.Valid {
			br.Description = description.String
		} else {
			br.Description = ""
		}
		if err != nil {
			log.Println(errors.Wrap(err, "GetAllRules Scanning"))
			return nil, err
		}
		rules = append(rules, br)
	}

	err = rows.Err()
	if err != nil {
		log.Println(errors.Wrap(err, "GetAllRules Row.Err"))
		return nil, err
	}

	rulesObject := models.RulesObject{Rules: rules, NumRules: len(rules)}
	return &rulesObject, nil
}

func (c *Client) GetSingleRule(id int) (*models.BounceRule, error) {
	var br models.BounceRule
	err := c.Conn.QueryRow("SELECT * From bounce_rule WHERE id = ?", id).Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
	if err != nil {
		log.Println(errors.Wrap(err, "GetSingleRule"))
		return nil, err
	}
	return &br, nil
}

func (c *Client) CreateRule(rule *models.BounceRule) error {
	_, err := c.Conn.Exec("INSERT INTO bounce_rule(response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?)", rule.ResponseCode, rule.EnhancedCode, rule.Regex, rule.Priority, rule.Description, rule.BounceAction)
	if err != nil {
		log.Println(errors.Wrap(err, "CreateRule"))
		return err
	}
	return nil
}

func (c *Client) UpdateRule(newRule *models.BounceRule) error {
	_, err := c.Conn.Exec("UPDATE bounce_rule SET id=? , response_code= ? , enhanced_code= ? , regex= ?, priority= ? , description= ?, bounce_action= ? WHERE id= ?", newRule.ID, newRule.ResponseCode, newRule.EnhancedCode, newRule.Regex, newRule.Priority, newRule.Description, newRule.BounceAction, newRule.ID)
	if err != nil {
		log.Println(errors.Wrap(err, "UpdateRule"))
		return err
	}
	return nil
}

func (c *Client) DeleteRule(id int) error {
	_, err := c.Conn.Exec("DELETE FROM bounce_rule WHERE id= ?", id)
	if err != nil {
		log.Println(errors.Wrap(err, "DeleteRule"))
		return err
	}
	return nil
}