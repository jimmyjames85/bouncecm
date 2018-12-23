package db

import (
	"log"
	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// Changelog - Function to pull all rules from db
func  (c *Client) Changelog() (*[]models.ChangelogEntry, error) {
	rules := []models.ChangelogEntry{}

	rows, err := c.Conn.Query("SELECT * FROM changelog ")


	if err != nil {
		log.Println(err)
		return nil, err;
	}

	defer rows.Close()

	for rows.Next() {
		br := models.ChangelogEntry{}

		err = rows.Scan(&br.ID,  &br.UserID,  &br.Comment,  &br.CreatedAt, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
		
		if err != nil {
			log.Println(err)
			return nil, err;
		}

		rules = append(rules, br)
	}

	err = rows.Err()
    if err != nil {
        return nil, err
	}
	
	return &rules, nil
}
