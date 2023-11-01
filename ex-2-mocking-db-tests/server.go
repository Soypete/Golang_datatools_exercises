package server

import (
	"net/http"
)

type Connector interface {
	selectUser(username string) (User, error)
	updateUser(user User) error
	deleteUser(username string) error
}

type Connection struct {
	db connector
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// endpoint: /getUsername/{username}
func (c *Connection) getUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/getUsername/"):]
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}
	user, err := c.db.selectUser(username)
	if err != nil {
		http.Error(w, http.StatusText(500)+", error while selecting user", 500)
		return
	}
	w.Status = 200
	w.Write([]byte(user.Name))

	// return username and 200
}

// endpoint: /updateUsername/{username}
func updateUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/updateUsername/"):]
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return response and 200
	w.Write([]byte("user registered"))
}

// endpoint: /deleteUsername/{username}
func deleteUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/deleteUsername/"):]
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return response and 200
	w.Write([]byte("user deleted")) // TODO: return deleted username
}
