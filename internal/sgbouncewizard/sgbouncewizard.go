package sgbouncewizard

import (
	"context"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/models"
	"github.com/jimmyjames85/bouncecm/internal/config"
	"golang.org/x/crypto/bcrypt"
	"github.com/pkg/errors"
)

type Server struct {
	DBClient *db.Client
}

// NewServer ...
func NewServer(c config.Configuration) (*Server, error) {
	client, err := db.NewDB(c)
	if err != nil {
		log.Println(errors.Wrap(err, "Connecting to Client Failed"))
		return nil, errors.Wrap(err, "Connecting to Client Failed")
	}

	err = client.Conn.Ping()

	if err != nil {
		log.Println(errors.Wrap(err, "Ping Failed"))
		return nil, errors.Wrap(err, "Ping Failed")
	}

	return &Server{DBClient: client}, nil
}

func (srv *Server) RuleContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rule *models.BounceRule

		bounce_id := chi.URLParam(r, "bounce_id"); 
		bouncd_idInt, err := strconv.Atoi(bounce_id)
		
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if bounce_id != "" {
			rule, err = srv.DBClient.GetSingleRule(bouncd_idInt)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		ctx := context.WithValue(r.Context(), "rule", rule)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (srv *Server) generateHash(pwd []byte) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(errors.Wrap(err, "generateHash"))
		return nil, err
	}
	result := string(hash)
	return &result, nil
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
	c := models.UserCredentials{}

	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := srv.DBClient.GetUserByEmail(c.Email)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// ListRules - wrapper to grab all rules
func  (srv *Server) getAllRulesRoute(w http.ResponseWriter, r *http.Request) {
	rules, err := srv.DBClient.GetAllRules()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&rules)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (srv *Server) getRuleRoute(w http.ResponseWriter, r *http.Request) {
	rule := r.Context().Value("rule").(*models.BounceRule)

	data, err := json.Marshal(rule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)  
}

func (srv *Server) deleteRuleRoute(w http.ResponseWriter, r *http.Request) {
	toDelete := r.Context().Value("rule").(*models.BounceRule)
	err := srv.DBClient.DeleteRule(toDelete.ID)
	
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}


func (srv *Server) createRuleRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rule models.BounceRule
	
	err := decoder.Decode(&rule)
	
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = srv.DBClient.CreateRule(&rule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (srv *Server) updateRuleRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newRule models.BounceRule
	err := decoder.Decode(&newRule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	srv.DBClient.UpdateRule(&newRule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}


func (srv *Server) GetChangelog(w http.ResponseWriter, r *http.Request) {
	rules, err := srv.DBClient.Changelog()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(*rules)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (srv *Server) Serve(Port int) {
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
		r.Get("/", srv.getAllRulesRoute)
		r.Post("/", srv.createRuleRoute)

		r.Route("/{bounce_id}", func(r chi.Router) {
			r.Use(srv.RuleContext)
			r.Get("/", srv.getRuleRoute)
			r.Delete("/", srv.deleteRuleRoute)
			r.Put("/", srv.updateRuleRoute)
		})
	})

	port := fmt.Sprintf(":%d", Port)
	http.ListenAndServe(port, r)

}