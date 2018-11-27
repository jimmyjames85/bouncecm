package db

import (
	"database/sql"
)

type Client struct {
	Conn *sql.DB
}

type Configuration struct {
}

// NewDB ...
func NewDB(c *Configuration) (*Client, error) {
	// user c to tweak your mysql setting
	// create a connection
	// return Client ref with conn
	conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/drop_rules")
	if err != nil {
		return nil, err
	}

	return &Client{Conn: conn}, nil
}
