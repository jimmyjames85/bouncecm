package db

import (
	"fmt"
	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// Changelog - Function to pull all rules from db
func  (c *Client) Changelog() (models.ChangelogTable, error) {
	rules := []models.ChangelogEntry{}

	rows, err := c.Conn.Query("SELECT * FROM changelog ")


	checkErr(err)

	for rows.Next() {
		br := models.ChangelogEntry{}

		err = rows.Scan(&br.ID,  &br.UserID,  &br.Comment,  &br.CreatedAt, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
		fmt.Println(br)
		checkErr(err)
		rules = append(rules, br)
	}

	ChangelogTable := models.ChangelogTable{Rules: rules, NumRules: len(rules)}

	return ChangelogTable, nil
}
