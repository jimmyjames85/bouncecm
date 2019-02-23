package sgbouncewizard

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jimmyjames85/bouncecm/internal/config"
	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
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

		bounce_id := chi.URLParam(r, "bounce_id")
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

	limit_param, ok := queryParams["limit"]

	if !ok || len(limit_param) > 1{
		paramError := errors.New("Invalid limit Parameter: does not exist or to many")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}


	offset_param, ok := queryParams["offset"]
	if !ok ||  len(offset_param) > 1 {
		paramError := errors.New("Invalid offset Parameter: does not exist or to many")
		http.Error(w, paramError.Error(), http.StatusBadRequest)
		return
	}


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

	rules, err := srv.DBClient.GetAllRules(Offset,limit)

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

		bounce_id := chi.URLParam(r, "bounce_id")

		bouncd_idInt, err := strconv.Atoi(bounce_id)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		queryParams := r.URL.Query()

		limit, ok := queryParams["limit"]

		if len(limit) > 1 {
			paramError := errors.New("Duplicate Parameters")
			http.Error(w, paramError.Error(), http.StatusBadRequest)
		}

		if !ok {
			changelog, err = srv.DBClient.GetChangeLogEntries(bouncd_idInt, nil)
			if err != nil {
				if strings.HasSuffix(err.Error(), "no rows in result set") {
					log.Println(err)
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
			}
		} else {

			limitAsInt, err := strconv.Atoi(r.FormValue("limit"))

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

	port := fmt.Sprintf(":%d", Port)
	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Println(errors.Wrap(err, "ListenaAndServer Failed"))
		return
	}
}
