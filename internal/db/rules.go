package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
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


func getRuleDB(id string) (*models.BounceRule, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	var bounce_rule *models.BounceRule
	checkErr(err)
	rows, err := db.Query("SELECT * FROM bounce_rule WHERE id=" + id)
	for rows.Next() {
		var br models.BounceRule
		err = rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
		checkErr(err)
		bounce_rule = &br
		return bounce_rule, nil
	}
	checkErr(err)
	defer rows.Close()
	defer db.Close()
	return nil, err
}
func deleteRuleDB(id int) (*models.BounceRule, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	checkErr(err)
	query := fmt.Sprintf("%s%d", "DELETE FROM bounce_rule WHERE id=", id)
	_, err = db.Query(query)
	checkErr(err)
	defer db.Close()
	return nil, err
}
func updateRuleDB(ruleDifferences map[string]interface{}, prevRule *models.BounceRule) {
	queryString := createUpdateQuery(ruleDifferences, prevRule)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	checkErr(err)
	_, err = db.Query(queryString)
	checkErr(err)
	db.Close()
}

func getRuleDifferences(prevRule *models.BounceRule, newRule *models.BounceRule) map[string]interface{} {
	ruleChange := make(map[string]interface{})
	prevRuleValue := reflect.ValueOf(prevRule).Elem()
	newRuleValue := reflect.ValueOf(newRule).Elem()
	bounceRuleStruct := prevRuleValue.Type()
	for i := 0; i < prevRuleValue.NumField(); i++ {
		prevRuleField := prevRuleValue.Field(i)
		newRuleField := newRuleValue.Field(i)
		if prevRuleField.Kind() == reflect.Ptr {
			if prevRuleField.Elem() != newRuleField.Elem() {
				ruleChange[strings.ToLower(bounceRuleStruct.Field(i).Name)] = newRuleField.Elem().String()
			}
		} else {
			if prevRuleField.Interface() != newRuleField.Interface() {
				ruleChange[strings.ToLower(bounceRuleStruct.Field(i).Name)] = newRuleField.Interface()
			}
		}
	}
	return ruleChange
}

func createUpdateQuery(ruleDifferences map[string]interface{}, prevRule *models.BounceRule) string {
	var queryString string = "UPDATE bounce_rule SET "
	innerCount := 1
	for k, v := range ruleDifferences {
		queryString += k + "="
		switch v.(type) {
		case string:
			queryString += ("\"" + v.(string) + "\"")
			break
		case int:
			queryString += ("\"" + strconv.Itoa(v.(int)) + "\"")
			break
		}
		if innerCount < len(ruleDifferences) {
			queryString += ","
		}
		innerCount++
	}
	queryString += (" WHERE id=" + strconv.Itoa(prevRule.ID))
	return queryString
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
