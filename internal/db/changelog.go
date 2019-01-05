package db

import (
	"github.com/pkg/errors"
	"time"
	"database/sql"
	"math"
	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// Changelog - Function to pull all rules from db
func  (c *Client) GetAllChangelogEntries() ([]models.BounceRule, error) {
	rules := []models.BounceRule{}

	rows, err := c.Conn.Query("SELECT * FROM changelog ")


	if err != nil {
		return nil, errors.Wrap(err, "Query Error")
	}

	defer rows.Close()

	for rows.Next() {
		br := models.BounceRule{}

		err = rows.Scan(&br.ID,  &br.UserID,  &br.Comment,  &br.CreatedAt, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
		
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



func (c *Client) GetChangeLogEntries(id int, limit *int) ([]models.BounceRule, error) {

	stmt, err := c.Conn.Prepare("SELECT * From changelog WHERE rule_id = ? LIMIT ?")


	var rows *sql.Rows


	if limit == nil{
		rows, err = stmt.Query(id, math.MaxInt64)
	} else {
		rows, err = stmt.Query(id , limit)
	}

	if err != nil {
		return nil, errors.Wrap(err, "GetChangeLogEntries Query")
	}
	
	rules := []models.BounceRule{}


	defer rows.Close()

	for rows.Next() {
		cl := models.BounceRule{}

		err = rows.Scan(&cl.ID,  &cl.UserID,  &cl.Comment,  &cl.CreatedAt, &cl.ResponseCode, &cl.EnhancedCode, &cl.Regex, &cl.Priority, &cl.Description, &cl.BounceAction)
		
		if err != nil {
			return nil, errors.Wrap(err, "GetChangeLogEntries Row Scan")
		}
		rules = append(rules, cl)
	}
	return rules, nil
}

func (c *Client) CreateChangeLogEntry(lastId int, entry *models.BounceRule) error {
	_, err := c.Conn.Exec("INSERT INTO changelog(rule_id,user_id,comment,created_at,response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?,?,?,?,?)", lastId, entry.UserID, entry.Comment,int32(time.Now().Unix()) ,entry.ResponseCode, entry.EnhancedCode, entry.Regex, entry.Priority, entry.Description, entry.BounceAction)
	if err != nil {
		return errors.Wrap(err, "CreateChangeLogEntry")
	}
	return nil
}


