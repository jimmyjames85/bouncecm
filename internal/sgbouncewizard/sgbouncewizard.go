package sgbouncewizard

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jimmyjames85/bouncecm/internal/config"
	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/olahol/melody.v1"
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

		bouncd_id, err := strconv.Atoi(chi.URLParam(r, "bounce_id"))

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rule, err = srv.DBClient.GetSingleRule(bouncd_id)
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
		return "", errors.Wrap(err, "failed to gen hash")
	}
	result := string(hash)
	return result, nil
}

func (srv *Server) verifyPassword(hashed string, plain []byte) error {
	byteHash := []byte(hashed)

	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	if err != nil {
		return err
	}
	return nil
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

	if len(user) == 0 {
		passError := errors.New("User not found")
		http.Error(w, passError.Error(), http.StatusUnauthorized)
		return
	}

	result := models.UserObject{}
	err = srv.verifyPassword(user[0].Hash, []byte(c.Password))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	result.ID = user[0].ID
	result.FirstName = user[0].FirstName
	result.LastName = user[0].LastName
	result.Role = user[0].Role

	data, err := json.Marshal(result)
	log.Println("No user")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// ListRules - wrapper to grab all rules
func (srv *Server) GetAllRulesRoute(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	limit_param := queryParams["limit"]

	if len(limit_param) > 1 {
		paramError := errors.New("Invalid limit Parameter: does not exist or to many")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}

	offset_param := queryParams["offset"]
	if len(offset_param) > 1 {
		paramError := errors.New("Invalid offset Parameter: does not exist or to many")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}

	var rules []models.BounceRule
	var err error

	if len(offset_param) == 0 && len(limit_param) == 0 {
		rules, err = srv.DBClient.GetAllRules()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	} else {
		limit, err := strconv.Atoi(r.FormValue("limit"))

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		Offset, err := strconv.Atoi(r.FormValue("offset"))

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rules, err = srv.DBClient.GetAllRulesLimited(Offset, limit)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

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

func (srv *Server) GetRuleRoute(w http.ResponseWriter, r *http.Request) {
	rule, ok := r.Context().Value("rule").(*models.BounceRule)

	if !ok {
		log.Println("ContextValue of rule in GetRuleRoute: " + strconv.FormatBool(ok))
		paramError := errors.New("Route Parameters")
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

func (srv *Server) DeleteRuleRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var toDelete models.ChangelogEntry
	err := decoder.Decode(&toDelete)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	toDelete.Operation = "Delete"

	err = srv.DBClient.CreateChangeLogEntry(toDelete.ID, &toDelete)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.DeleteRule(toDelete.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (srv *Server) CreateRuleRoute(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var rule models.ChangelogEntry

	err := decoder.Decode(&rule)
	rule.Operation = "Create"
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	LastInsertedID, err := srv.DBClient.CreateRule(&rule.BounceRule)
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

	newRuleID, err := json.Marshal(TempJsonObject{"id": LastInsertedID})

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newRuleID)
}

func (srv *Server) UpdateRuleRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newRule models.ChangelogEntry
	err := decoder.Decode(&newRule)

	newRule.Operation = "Update"
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.UpdateRule(&newRule.BounceRule)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
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
		var changelog []models.ChangelogEntry

		bounce_id, err := strconv.Atoi(chi.URLParam(r, "bounce_id"))

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		queryParams := r.URL.Query()

		limit_param := queryParams["limit"]
		if len(limit_param) > 1 {
			paramError := errors.New("Invalid limit Parameter: does not exist or to many")
			http.Error(w, paramError.Error(), http.StatusBadRequest)
			return
		}

		offset_param := queryParams["offset"]
		if len(offset_param) > 1 {
			paramError := errors.New("Invalid offset Parameter: does not exist or to many")
			http.Error(w, paramError.Error(), http.StatusBadRequest)
			return
		}

		if len(offset_param) == 0 && len(limit_param) == 0 {
			changelog, err = srv.DBClient.GetChangeLogById(bounce_id)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		} else {
			limit, err := strconv.Atoi(r.FormValue("limit"))

			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			offset, err := strconv.Atoi(r.FormValue("offset"))

			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			changelog, err = srv.DBClient.GetChangeLogByIdLimited(bounce_id, offset, limit)

			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		ctx := context.WithValue(r.Context(), "changelog", changelog)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (srv *Server) GetChangeLogEntriesRoute(w http.ResponseWriter, r *http.Request) {
	changelog, ok := r.Context().Value("changelog").([]models.ChangelogEntry)

	if !ok {
		log.Println("ContextValue of rule in GetChangeLogEntriesRoute: " + strconv.FormatBool(ok))
		paramError := errors.New("Route Parameters")
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

	queryParams := r.URL.Query()
	limit_param := queryParams["limit"]

	if len(limit_param) > 1 {
		paramError := errors.New("Invalid limit Parameter: Too many")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}

	offset_param := queryParams["offset"]

	if len(offset_param) > 1 {
		paramError := errors.New("Invalid offset Parameter: Too many")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}

	var changelog []models.ChangelogEntry
	var err error

	if len(offset_param) == 0 && len(limit_param) == 0 {

		changelog, err = srv.DBClient.GetAllChangelogEntries()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	} else {
		limit, err := strconv.Atoi(r.FormValue("limit"))

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		Offset, err := strconv.Atoi(r.FormValue("offset"))

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		changelog, err = srv.DBClient.GetAllChangelogEntriesLimited(Offset, limit)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
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
	var changelog models.ChangelogEntry

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (srv *Server) Serve(Port int) {
	r := chi.NewRouter()
	m := melody.New()
	lock := new(sync.Mutex)
	currentlyViewing := map[string]bool{}

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

	r.Route("/change_logs", func(r chi.Router) {
		r.Get("/", srv.GetAllChangelogEntries)
		r.Post("/", srv.createChangeLogEntryRoute)

		r.Route("/{bounce_id}", func(r chi.Router) {
			r.Use(srv.ChangelogContext)
			r.Get("/", srv.GetChangeLogEntriesRoute)

		})
	})

	r.Route("/bounce_rules", func(r chi.Router) {
		r.Get("/", srv.GetAllRulesRoute)
		r.Post("/", srv.CreateRuleRoute)

		r.Route("/{bounce_id}", func(r chi.Router) {
			r.Use(srv.RuleContext)
			r.Get("/", srv.GetRuleRoute)
			r.Delete("/", srv.DeleteRuleRoute)
			r.Put("/", srv.UpdateRuleRoute)
		})
	})

	r.Route("/ws", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			m.HandleRequest(w, r)
		})
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println(string(msg))
		params := strings.Split(string(msg), " ")
		lock.Lock()
		if params[0] == "edit" {
			if currentlyViewing[params[1]] {
				_, exists := s.Get(params[1])
				if !exists {
					log.Println("ATTEMPT TO EDIT RULE THAT IS BEING EDITED")
					s.Write([]byte("INUSE"))
				} else {
					s.Write([]byte("ALREADY"))
				}
			} else {
				s.Set(params[1], true)
				currentlyViewing[params[1]] = true
				s.Write([]byte("EDIT"))
				m.BroadcastOthers([]byte("INUSE"), s)
			}
		} else if params[0] == "release" {
			if currentlyViewing[params[1]] {
				delete(currentlyViewing, params[1])
				s.Set(params[1], false)
				m.BroadcastOthers([]byte("FREE"), s)
			}
		} else if params[0] == "check" {
			if currentlyViewing[params[1]] {
				_, exists := s.Get(params[1])
				if !exists {
					s.Write([]byte("INUSE"))
				}
			} else {
				s.Write([]byte("FREE"))
			}
		}
		lock.Unlock()
	})

	port := fmt.Sprintf(":%d", Port)
	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Println(errors.Wrap(err, "ListenaAndServer Failed"))
		return
	}
}
