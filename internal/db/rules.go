package db

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"encoding/json"

	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// ListRules - Function to pull all rules from db
func (c *Client) ListRules() (models.RulesObject, error) {
	rules := []models.BounceRule{}

	rows, err := c.Conn.Query("SELECT * FROM bounce_rule")

	checkErr(err)

	for rows.Next() {
		br := models.BounceRule{}

		err = rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
		checkErr(err)
		rules = append(rules, br)
	}

	rulesObject := models.RulesObject{Rules: rules, NumRules: len(rules)}

	return rulesObject, nil
}

func (c *Client) CreateRuleDB(rule *models.BounceRule) []byte {
	stmt, err := c.Conn.Prepare("INSERT INTO bounce_rule(id,response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Query(rule.ID, rule.ResponseCode, rule.EnhancedCode, rule.Regex, rule.Priority, rule.Description, rule.BounceAction)
	checkErr(err)

	data, err := json.Marshal(rule)
	checkErr(err)
	return data
}

func (c *Client) GetRuleDB(id string) (*models.BounceRule, error) {
	var bounce_rule *models.BounceRule
	rows, err := c.Conn.Query("SELECT * FROM bounce_rule WHERE id=" + id)
	for rows.Next() {
		var br models.BounceRule
		err = rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
		checkErr(err)
		bounce_rule = &br
		return bounce_rule, nil
	}
	checkErr(err)
	return nil, err
}

func (c *Client) DeleteRuleDB(id int) {
	query := fmt.Sprintf("%s%d", "DELETE FROM bounce_rule WHERE id=", id)
	_, err := c.Conn.Query(query)
	checkErr(err)
}

func (c *Client) UpdateRuleDB(ruleDifferences map[string]interface{}, prevRule *models.BounceRule) {
	queryString := createUpdateQuery(ruleDifferences, prevRule)
	_, err :=  c.Conn.Query(queryString)
	checkErr(err)
}

// getRuleDifferences this and the next functions don't require the the DB client to function do they need
// need the (c *Client)?
func GetRuleDifferences(prevRule *models.BounceRule, newRule *models.BounceRule) (map[string]interface{}) {
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

func  createUpdateQuery(ruleDifferences map[string]interface{}, prevRule *models.BounceRule) string {
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
