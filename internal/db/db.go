package db

import (
	"database/sql"
	"github.com/jimmyjames85/bouncecm/internal/config"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type Client struct {
	Conn *sql.DB
}

// NewDB ... 
func NewDB(c config.Configuration) (*Client, error) {
	fmt.Println(c)
	dbConf := &mysql.Config{
		User:                 c.DBUser,
		Passwd:               c.DBPass,
		Addr:                 fmt.Sprintf("%s:%d", c.DBHost, c.DBPort),
		Net:                  "tcp",
		ReadTimeout:          c.DBReadTimeout,
		WriteTimeout:         c.DBWriteTimeout,
		AllowNativePasswords: true,
		DBName:				  c.DBName,
	}

	fmt.Println(dbConf.FormatDSN())

	// user c to tweak your mysql setting
	// create a connection
	// return Client ref with conn
	conn, err := sql.Open("mysql", dbConf.FormatDSN())
	if err != nil {
		return nil, err
	}

	return &Client{Conn: conn}, nil
}