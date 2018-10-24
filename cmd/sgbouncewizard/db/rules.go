package db

import (
	"../models"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func ListRules() (models.RulesObject, error) {
		rules := []models.BounceRule{};

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	
		checkErr(err);
		
		rows, err := db.Query("SELECT * FROM bounce_rule");
	
		checkErr(err);
	
		for rows.Next() {
			br := models.BounceRule{}
	
			err = rows.Scan(&br.Id, &br.Response_code, &br.Enhanced_code, &br.Regex, &br.Priority, &br.Description, &br.Bounce_action)
			checkErr(err)
			rules = append(rules, br);
		}
	
		defer rows.Close();
	
		db.Close();

		rulesObject := models.RulesObject{Rules: rules, NumRules: len(rules)}

		return rulesObject, nil;
  }

  func checkErr(err error) {
	if err != nil {
		panic(err)
	}
	}