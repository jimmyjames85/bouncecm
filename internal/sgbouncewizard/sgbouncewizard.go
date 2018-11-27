package sgbouncewizard

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jimmyjames85/bouncecm/internal/db"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

type Server struct {
	DBClient *db.Client
}

// NewServer ...
func NewServer(c *db.Configuration) (*Server, error) {
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
		r.Post("/", CheckUser)
	})

	r.Route("/changelogs", func(r chi.Router) {
		r.Get("/", GetChangelog)
	})

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

	http.ListenAndServe(":3000", r)

}
