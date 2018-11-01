package db

import (
	"database/sql"
	"fmt"

	"../models"
	_ "github.com/go-sql-driver/mysql"
)

// ListRules - Function to pull all rules from db
func ListRules() (models.RulesObject, error) {
	rules := []models.BounceRule{}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")

	checkErr(err)

	rows, err := db.Query("SELECT * FROM bounce_rule")

	checkErr(err)

	for rows.Next() {
		br := models.BounceRule{}

		err = rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
		checkErr(err)
		rules = append(rules, br)
	}

	defer rows.Close()

	db.Close()

	rulesObject := models.RulesObject{Rules: rules, NumRules: len(rules)}

	return rulesObject, nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
