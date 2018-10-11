package main

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type BounceRule struct {
	Id int `json:"id"`;
	Response_code string `json:"response_code"`;
	Enhanced_code string `json:"enhanced_code"`;
	Regex string `json:"regex"`;
	Priority int `json:"priority"`;
	Description []byte `json:"description"`;
	Bounce_action string `json:"bounce_action"`;
}

type RulesObject struct {
	Rules []BounceRule `json:"rules"`;
	NumRules int `json:"numRules"`
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/bounce_rules", func(r chi.Router) {
		r.Get("/", listRules);
	})

	http.ListenAndServe(":3000", r)

}

func listRules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
    } else {
		var rules []BounceRule;

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	
		checkErr(err);
		
		rows, err := db.Query("SELECT * FROM bounce_rule");
	
		checkErr(err);
	
		for rows.Next() {
			var br BounceRule
	
			err = rows.Scan(&br.Id, &br.Response_code, &br.Enhanced_code, &br.Regex, &br.Priority, &br.Description, &br.Bounce_action)
			checkErr(err)
			rules = append(rules, br);
		}
	
		defer rows.Close();
	
		db.Close();

		rulesObject := RulesObject{Rules: rules, NumRules: len(rules)}
	
		data, err := json.Marshal(rulesObject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
    }
  }

  func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}