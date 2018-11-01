package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jimmyjames85/bouncecm/internal/models"
	"github.com/go-sql-driver/mysql"
)

var _ = mysql.Open

// GetUserByEmail - Function to pull user from db
func GetUserByEmail(email string) ([]*models.User, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	fmt.Println(email)
	rows, err := db.Query("SELECT * FROM user where email = ?", email)

	if err != nil {
		return nil, errors.New("User not found")
	}

	result := []*models.User{}

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
