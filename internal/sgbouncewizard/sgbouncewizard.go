package sgbouncewizard

import (
	"context"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"fmt"
	"strings"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/models"
	"github.com/jimmyjames85/bouncecm/internal/config"
	"golang.org/x/crypto/bcrypt"
	"github.com/pkg/errors"
	"github.com/go-chi/cors"
)

type TempJsonObject map[string]interface{}

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

	
		rule, err = srv.DBClient.GetSingleRule(bouncd_idInt)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "rule", rule)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (srv *Server) generateHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(errors.Wrap(err, "generateHash"))
		return "", errors.Wrap(err,"failed to gen hash")
	}
	result := string(hash)
	return result, nil
}


func (srv *Server) verifyPassword(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)

	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	if err != nil {
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
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	result := models.UserObject{}
	if srv.verifyPassword(user[0].Hash, []byte(c.Password)) {
		result.ID = user[0].ID
		result.FirstName = user[0].FirstName
		result.LastName = user[0].LastName
		result.Role = user[0].Role
	} else {
		passError :=  errors.New("verifyPassword Failed")
		http.Error(w, passError.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}


func  (srv *Server) getAllRulesRoute(w http.ResponseWriter, r *http.Request) {
	rules, err := srv.DBClient.GetAllRules()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(&rules)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (srv *Server) getRuleRoute(w http.ResponseWriter, r *http.Request) {
	rule , ok := r.Context().Value("rule").(*models.BounceRule)

	if !ok {
		log.Println("ContextValue of rule in GetRuleRoute: " + strconv.FormatBool(ok))
		paramError :=  errors.New("Route Parameters")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(rule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)  
}

func (srv *Server) deleteRuleRoute(w http.ResponseWriter, r *http.Request) {
	toDelete, ok := r.Context().Value("rule").(*models.BounceRule)

	if !ok {
		log.Println("ContextValue of rule in GetRuleRoute: " + strconv.FormatBool(ok))
		paramError :=  errors.New("Route Parameters")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}

	err := srv.DBClient.DeleteRule(toDelete.ID)
	
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}


func (srv *Server) createRuleRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rule models.BounceRule
	
	err := decoder.Decode(&rule)
	
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	LastInsertedID, err := srv.DBClient.CreateRule(&rule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.CreateChangeLogEntry(LastInsertedID, &rule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newRuleID , err := json.Marshal(TempJsonObject{"id": LastInsertedID})

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newRuleID)
}

func (srv *Server) updateRuleRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newRule models.BounceRule
	err := decoder.Decode(&newRule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.UpdateRule(&newRule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.CreateChangeLogEntry(newRule.ID, &newRule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}


func (srv *Server) ChangelogContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var changelog []models.BounceRule

		bounce_id := chi.URLParam(r, "bounce_id"); 

		
		
	
		bouncd_idInt, err := strconv.Atoi(bounce_id)
		
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}


		queryParams := r.URL.Query()

		limit, ok := queryParams["limit"];
	
		if  !ok {
			changelog, err = srv.DBClient.GetChangeLogEntries(bouncd_idInt, nil)
		} else {
			
			limitAsString := strings.Join(limit, "")
			limitAsInt, err := strconv.Atoi(limitAsString)

			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			
			changelog, err = srv.DBClient.GetChangeLogEntries(bouncd_idInt, &limitAsInt)

			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

		}




		ctx := context.WithValue(r.Context(), "changelog", changelog)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (srv *Server) GetChangeLogEntriesRoute(w http.ResponseWriter, r *http.Request) {
	changelog , ok := r.Context().Value("changelog").([]models.BounceRule)

	if !ok {
		log.Println("ContextValue of rule in GetChangeLogEntriesRoute: " + strconv.FormatBool(ok))
		paramError :=  errors.New("Route Parameters")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(changelog)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (srv *Server) GetAllChangelogEntries(w http.ResponseWriter, r *http.Request) {
	changelog, err := srv.DBClient.GetAllChangelogEntries()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(changelog)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// createChangeLogEntryRoute will only be used to seed the change log databse 
// with all current rules
func (srv *Server) createChangeLogEntryRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var changelog models.BounceRule

	err := decoder.Decode(&changelog)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.CreateChangeLogEntry(changelog.ID, &changelog)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}


func (srv *Server) Serve(Port int) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome: Routes \n\n /bounce_rules/{id} - Get Post Put Delete \n\n /bounce_rules - Get \n\n /change_log - Get \n\n /change_log/{id} - Get"))
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/", srv.CheckUser)
	})

	r.Route("/changelogs", func(r chi.Router) {
		r.Get("/", srv.GetAllChangelogEntries)
		r.Post("/", srv.createChangeLogEntryRoute)

		r.Route("/{bounce_id}", func(r chi.Router) {
			r.Use(srv.ChangelogContext)
			r.Get("/", srv.GetChangeLogEntriesRoute)

		})
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
	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Println(errors.Wrap(err, "ListenaAndServer Failed"))
		return
	}
}