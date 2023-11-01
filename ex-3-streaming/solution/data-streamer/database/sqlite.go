package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Client struct {
	db *sql.DB
}

func Setup() (*Client, error) {
	fileName := "database/ex-1-connection/sqlite.db"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	return &Client{db: db}, nil
}

func (c *Client) Close() error {
	return c.db.Close()
}

// GetUser gets a user from the database by ID.
func (c *Client) GetUser(ctx context.Context, id int) (User, error) {
	var user User
	err := c.db.QueryRow("SELECT * FROM users WHERE id = $", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return User{}, fmt.Errorf("error getting user: %v", err)
	}
	return user, nil
}
