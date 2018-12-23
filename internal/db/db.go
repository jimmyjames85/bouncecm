package db

import (
	"database/sql"
	"github.com/jimmyjames85/bouncecm/internal/config"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql"
)

type Client struct {
	Conn *sql.DB
}

// NewDB ... 
func NewDB(c config.Configuration) (*Client, error) {
	dbConf := &mysql.Config{
		User:                 c.DBUser,
		Passwd:               c.DBPass,
		Addr:                 fmt.Sprintf("%s:%d", c.DBHost, c.DBPort),
		Net:                  "tcp",
		ReadTimeout:          c.DBReadTimeout,
		WriteTimeout:         c.DBWriteTimeout,
		AllowNativePasswords: true,
		DBName:            	  c.DBName,
	}

	log.Println(dbConf.FormatDSN())

	conn, err := sql.Open("mysql", dbConf.FormatDSN())
	if err != nil {
		return nil, err
	}

	return &Client{Conn: conn}, nil
}

// like this?
func  (c *Client) Ping() (error) {
	return nil
}
