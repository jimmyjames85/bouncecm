package db

import (
	"github.com/pkg/errors"
	// Blank import required for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmyjames85/bouncecm/internal/models"
)

// GetUserByEmail - Function to pull user from db
func (c *Client) GetUserByEmail(email string) ([]*models.User, error) {
	rows, err := c.Conn.Query("SELECT * FROM user where email = ?", email)

	if err != nil {
		return nil, errors.Wrap(err, "GetUserByEmail User Not Found")
	}

	result := []*models.User{}

	for rows.Next() {
		u := models.User{}

		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Role, &u.Hash, &u.CreatedAt)

		if err != nil {
			return nil, errors.Wrap(err, "GetUserByEmail Cannot add to list")
		}

		result = append(result, &u)
	}

	err = rows.Err()
    if err != nil {
        return nil, errors.Wrap(err, "GetUserByEmail Row Error")
	}

	return result, nil
}
