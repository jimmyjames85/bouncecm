package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"./db"
	"./models"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
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

	r.Route("/bounce_rules", func(r chi.Router) {
		r.Get("/", ListRules);
	})

	http.ListenAndServe(":3000", r)

}


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
	}

	return string(hash)
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
		result.Id = user[0].Id;
		result.First_name = user[0].First_name
		result.Last_name = user[0].Last_name
		result.Role = user[0].Role
	}

	data, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

