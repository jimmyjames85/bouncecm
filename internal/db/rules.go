package db

import (
	"fmt"

	"log"
	"strconv"

	"encoding/json"

	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// ListRules - Function to pull all rules from db
func (c *Client) ListRules() (*models.RulesObject, error) {
	rules := []models.BounceRule{}

	rows, err := c.Conn.Query("SELECT * FROM bounce_rule")

	if err != nil {
		log.Println(err)
		fmt.Println("FAILEDHERE?")

		return nil, err;

	}

	for rows.Next() {
		br := models.BounceRule{}
		var description *string
		err := rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &description , &br.BounceAction)
		
		// case when description is null becuase it can be null according to droprules.sql
		if description == nil {
			br.Description = ""
		} else {
			br.Description = *description
		}


		if err != nil {
			log.Println(err)
			fmt.Println(br)

			return nil, err;
		}

		rules = append(rules, br)
		
	}

	rulesObject := models.RulesObject{Rules: rules, NumRules: len(rules)}

	return &rulesObject, nil
}

func (c *Client) CreateRuleDB(rule *models.BounceRule) ([]byte, error) {
	stmt, err := c.Conn.Prepare("INSERT INTO bounce_rule(response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?)")
	
	if err != nil {
		log.Println(err)
		return nil, err;
	}


	_, err = stmt.Exec(rule.ResponseCode, rule.EnhancedCode, rule.Regex, rule.Priority, rule.Description, rule.BounceAction)
	
	if err != nil {
		log.Println(err)
		return nil, err;
	}

	data, err := json.Marshal(rule)
	
	if err != nil {
		log.Println(err)
		return nil, err;
	}
	return data, nil
}

func (c *Client) GetRuleDB(id string) (*models.BounceRule, error) {
	var bounce_rule *models.BounceRule
	rows, err := c.Conn.Query("SELECT * FROM bounce_rule WHERE id=" + id)
	for rows.Next() {
		var br models.BounceRule
		err = rows.Scan(&br.ID, &br.ResponseCode, &br.EnhancedCode, &br.Regex, &br.Priority, &br.Description, &br.BounceAction)
		if err != nil {
			log.Println(err)
			return nil, err;
		}
		bounce_rule = &br
		return bounce_rule, nil
	}
	return nil, err
}

func (c *Client) DeleteRuleDB(id int) (error){
	query := fmt.Sprintf("%s%d", "DELETE FROM bounce_rule WHERE id=", id)
	_, err := c.Conn.Query(query)

	if err != nil {
		log.Println(err)
		return err;
	}
	return nil;
}

func (c *Client) UpdateRuleDB(ruleDifferences map[string]interface{}, prevRule *models.BounceRule) error {
	queryString := createUpdateQuery(ruleDifferences, prevRule)
	fmt.Println(queryString)
	_, err :=  c.Conn.Query(queryString)
	if err != nil {
		fmt.Println("fails here")

		log.Println(err)
		return err;
	}
	return nil;
}

// getRuleDifferences this and the next functions don't require the the DB client to function do they need
// need the (c *Client)?
func GetRuleDifferences(prevRule *models.BounceRule, newRule *models.BounceRule) (map[string]interface{}) {
	ruleChange := make(map[string]interface{})

	if (prevRule.ID != newRule.ID){
		ruleChange["id"] = newRule.ID
	}
	if (prevRule.ResponseCode != newRule.ResponseCode){
		ruleChange["response_code"] = newRule.ResponseCode
	}
	if (prevRule.EnhancedCode != newRule.EnhancedCode){
		ruleChange["enhanced_code"] = newRule.EnhancedCode
	}
	if (prevRule.Regex != newRule.Regex){
		ruleChange["regex"] = newRule.Regex
	}
	if (prevRule.Priority != newRule.Priority){
		ruleChange["priority"] = newRule.Priority
	}
	if (prevRule.Description != newRule.Description){
		ruleChange["description"] = newRule.Description
	}
	if (prevRule.BounceAction != newRule.BounceAction){
		ruleChange["bounce_action"] = newRule.BounceAction
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

