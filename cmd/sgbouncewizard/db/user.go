package db

import (
	"errors"
	"../models"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func GetUserByEmail(email string) ([]*models.User, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	fmt.Println(email)
	rows, err := db.Query("SELECT * FROM user where email = ?", email);

	if err != nil {
		return nil, errors.New("User not found")
	}

	result := []*models.User{}

	for rows.Next() {
		u := models.User{}

		err = rows.Scan(&u.Id, &u.First_name, &u.Last_name, &u.Email, &u.Role, &u.Hash, &u.Created_at)
		
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Cannot add to list")
		}

		result = append(result, &u);
	}

	return result, nil;
}