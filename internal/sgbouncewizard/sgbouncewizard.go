package sgbouncewizard

import (
	"context"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/models"
	"github.com/jimmyjames85/bouncecm/internal/config"

	"golang.org/x/crypto/bcrypt"

)

type Server struct {
	DBClient *db.Client
}

// NewServer ...
func NewServer(c config.Configuration) (*Server, error) {
	client, err := db.NewDB(c)
	if err != nil {
		return nil, err
	}

	return &Server{DBClient: client}, nil
}

func (srv *Server) RuleContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rule *models.BounceRule
		var err error
		if bounce_id := chi.URLParam(r, "bounce_id"); bounce_id != "" {
			rule, err = srv.DBClient.GetRuleDB(bounce_id)
			if err != nil {
				// do something to handle this error
			}
		}
		ctx := context.WithValue(r.Context(), "rule", rule)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (srv *Server) generateHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}


func (srv *Server) verifyPassword(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)

	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// CheckUser - wrapper function to auth user
func (srv *Server) CheckUser(w http.ResponseWriter, r *http.Request) {
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
	if srv.verifyPassword(user[0].Hash, []byte(c.Password)) {
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

// ListRules - wrapper to grab all rules
func  (srv *Server) ListRules(w http.ResponseWriter, r *http.Request) {
	rules, err := srv.DBClient.ListRules()

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

func (srv *Server) getRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rule := r.Context().Value("rule").(*models.BounceRule)
	data, err := json.Marshal(rule)
	checkErr(err)
	w.Write(data)  
}

func (srv *Server) deleteRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	toDelete := r.Context().Value("rule").(*models.BounceRule)
    srv.DBClient.DeleteRuleDB(toDelete.ID)
	data, err := json.Marshal(toDelete)
	checkErr(err)
	w.Write(data)
}


func (srv *Server) createRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var rule models.BounceRule
	err := decoder.Decode(&rule)
	checkErr(err)
	data := srv.DBClient.CreateRuleDB(&rule)
	w.Write(data)
}

func (srv *Server) updateRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prevRule := *r.Context().Value("rule").(*models.BounceRule)
	decoder := json.NewDecoder(r.Body)
	var newRule models.BounceRule
	err := decoder.Decode(&newRule)
	checkErr(err)
	data, err := json.Marshal(newRule)
	ruleDifferences := db.GetRuleDifferences(&prevRule, &newRule)
	srv.DBClient.UpdateRuleDB(ruleDifferences, &prevRule)
	w.Write(data)
}


func (srv *Server) GetChangelog(w http.ResponseWriter, r *http.Request) {
	rules, err := srv.DBClient.Changelog()

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

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (srv *Server) Serve() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/", srv.CheckUser)
	})

	r.Route("/changelogs", func(r chi.Router) {
		r.Get("/", srv.GetChangelog)
	})

	r.Route("/bounce_rules", func(r chi.Router) {
		r.Get("/", srv.ListRules)
		r.Post("/", srv.createRule)

		r.Route("/{bounce_id}", func(r chi.Router) {
			r.Use(srv.RuleContext)
			r.Get("/", srv.getRule)
			r.Delete("/", srv.deleteRule)
			r.Put("/", srv.updateRule)
		})
	})

	http.ListenAndServe(":3000", r)

}







// func checkErr(err error) {
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
	
