package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type BounceRule struct {
	Id            int     `json:"id"`
	Response_code string  `json:"response_code"`
	Enhanced_code string  `json:"enhanced_code"`
	Regex         string  `json:"regex"`
	Priority      int     `json:"priority"`
	Description   *string `json:"description"`
	Bounce_action string  `json:"bounce_action"`
}

type RulesObject struct {
	Rules    []BounceRule `json:"rules"`
	NumRules int          `json:"numRules"`
}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/bounce_rules", func(r chi.Router) {
		r.Get("/", ListRules)
		r.Post("/", createRule)

		r.Route("/{bounce_id}", func(r chi.Router) {
			r.Use(RuleContext)
			r.Get("/", getRule)
			r.Delete("/", deleteRule)
			r.Put("/", updateRule)
		})
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/", CheckUser)
	})

	http.ListenAndServe(":3000", r)
}

func RuleContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rule *BounceRule
		var err error
		if bounce_id := chi.URLParam(r, "bounce_id"); bounce_id != "" {
			rule, err = getRuleDB(bounce_id)
			checkErr(err)
		}
		ctx := context.WithValue(r.Context(), "rule", rule)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rule := r.Context().Value("rule").(*BounceRule)
	data, err := json.Marshal(rule)
	checkErr(err)
	w.Write(data)
}

func deleteRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	toDelete := r.Context().Value("rule").(*BounceRule)
	deleteRuleDB(toDelete.Id)
	data, err := json.Marshal(toDelete)
	checkErr(err)
	w.Write(data)
}

func createRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var rule BounceRule
	err := decoder.Decode(&rule)
	checkErr(err)
	data := createRuleDB(&rule)
	w.Write(data)
}

func updateRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prevRule := *r.Context().Value("rule").(*BounceRule)
	decoder := json.NewDecoder(r.Body)
	var newRule BounceRule
	err := decoder.Decode(&newRule)
	checkErr(err)

	data, err := json.Marshal(newRule)
	ruleDifferences := getRuleDifferences(&prevRule, &newRule)
	updateRuleDB(ruleDifferences, &prevRule)
	w.Write(data)
}

// ListRules - wrapper to grab all rules
func ListRules(w http.ResponseWriter, r *http.Request) {
	rules, err := db.ListRules()

	data, err := json.Marshal(rules)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func generateHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "error"
	}
	return string(hash)
}

func createRuleDB(rule *BounceRule) []byte {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO bounce_rule(id,response_code,enhanced_code,regex,priority,description,bounce_action) VALUES(?,?,?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Query(rule.Id, rule.Response_code, rule.Enhanced_code, rule.Regex, rule.Priority, rule.Description, rule.Bounce_action)
	checkErr(err)

	defer stmt.Close()
	defer db.Close()

	data, err := json.Marshal(rule)
	checkErr(err)
	return data
}

func getRuleDB(id string) (*BounceRule, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	var bounce_rule *BounceRule

	checkErr(err)
	rows, err := db.Query("SELECT * FROM bounce_rule WHERE id=" + id)

	for rows.Next() {
		var br BounceRule
		err = rows.Scan(&br.Id, &br.Response_code, &br.Enhanced_code, &br.Regex, &br.Priority, &br.Description, &br.Bounce_action)
		checkErr(err)
		bounce_rule = &br
		return bounce_rule, nil
	}
	checkErr(err)

	defer rows.Close()
	defer db.Close()
	return nil, err
}

func deleteRuleDB(id int) (*BounceRule, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	checkErr(err)
	query := fmt.Sprintf("%s%d", "DELETE FROM bounce_rule WHERE id=", id)
	_, err = db.Query(query)
	checkErr(err)
	defer db.Close()
	return nil, err
}

func updateRuleDB(ruleDifferences map[string]interface{}, prevRule *BounceRule) {
	queryString := createUpdateQuery(ruleDifferences, prevRule)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	checkErr(err)
	_, err = db.Query(queryString)
	checkErr(err)
	db.Close()
}

// Helper functions

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getRuleDifferences(prevRule *BounceRule, newRule *BounceRule) map[string]interface{} {
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

func createUpdateQuery(ruleDifferences map[string]interface{}, prevRule *BounceRule) string {
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
	queryString += (" WHERE id=" + strconv.Itoa(prevRule.Id))
	return queryString
}

func verifyPassword(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)

	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// CheckUser - wrapper function to auth user
func CheckUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	c := models.UserCredentials{}

	err := json.NewDecoder(r.Body).Decode(&c)

	user, err := db.GetUserByEmail(c.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := models.UserObject{}
	if verifyPassword(user[0].Hash, []byte(c.Password)) {
		result.ID = user[0].ID
		result.FirstName = user[0].FirstName
		result.LastName = user[0].LastName
		result.Role = user[0].Role
	}

	data, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
