package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// GetUserByEmail - Function to pull user from db
func GetUserByEmail(email string) ([]*models.User, error) {
	host := os.Getenv("DB_HOST")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")

	log.Println(host, pass, port, user)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/drop_rules", user, pass, host, port))

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM user where email = ?", email)

	if err != nil {
		return nil, errors.New("User not found")
	}

	result := []*models.User{}

	log.Println("get data")

	for rows.Next() {
		u := models.User{}

		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Role, &u.Hash, &u.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Cannot add to list")
		}

		result = append(result, &u)
	}

	return result, nil
}
