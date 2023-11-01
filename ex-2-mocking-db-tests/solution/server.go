package server

import (
	"encoding/json"
	"net/http"
)

type Connector interface {
	selectUser(email string) (User, error)
	updateUser(user User) error
	deleteUser(email string) error
}

type Connection struct {
	db Connector
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// endpoint: /getUser/{email}
func (c *Connection) getUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Path[len("/getUser/"):]
	// check if email is empty -> return 400
	if email == "" {
		http.Error(w, http.StatusText(400)+", email parameter cannot be empty", 400)
		return
	}
	user, err := c.db.selectUser(email)
	if err != nil {
		http.Error(w, http.StatusText(500)+", error while selecting user", 500)
		return
	}
	// return email and 200
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, http.StatusText(500)+", error while encoding user", 500)
		return
	}
}

// endpoint: /updateUser/{email}
func (c *Connection) updateUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Path[len("/updateUser/"):]
	// check if email is empty -> return 400
	if email == "" {
		http.Error(w, http.StatusText(400)+", email parameter cannot be empty", 400)
		return
	}
	defer r.Body.Close()
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, http.StatusText(500)+", error while decoding user", 500)
		return
	}

	err = c.db.updateUser(user)
	if err != nil {
		http.Error(w, http.StatusText(500)+", error while updating user", 500)
		return
	}

	// return response and 200
	w.WriteHeader(200)
	w.Write([]byte("user registered"))
}

// endpoint: /deleteUser/{email}
func (c *Connection) deleteUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Path[len("/deleteUser/"):]
	// check if email is empty -> return 400
	if email == "" {
		http.Error(w, http.StatusText(400)+", email parameter cannot be empty", 400)
		return
	}

	err := c.db.deleteUser(email)
	if err != nil {
		http.Error(w, http.StatusText(500)+", error while deleting user", 500)
		return
	}
	// return response and 200
	w.WriteHeader(200)
	w.Write([]byte("user deleted")) // TODO: return deleted email
}
