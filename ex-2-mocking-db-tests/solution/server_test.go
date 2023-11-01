package server

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

type dbPassMock struct{}

func (d *dbPassMock) selectUser(email string) (User, error) {
	return User{
		ID:    1,
		Name:  "Miriah Peterson",
		Email: "captainnobody1@email.com",
	}, nil
}

func (d *dbPassMock) updateUser(user User) error {
	return nil
}

func (d *dbPassMock) deleteUser(email string) error {
	return nil
}

type dbFailMock struct{}

func (d *dbFailMock) selectUser(email string) (User, error) {
	return User{}, fmt.Errorf("Error")
}

func (d *dbFailMock) updateUser(user User) error {
	return fmt.Errorf("Error")
}

func (d *dbFailMock) deleteUser(email string) error {
	return fmt.Errorf("Error")
}

func TestUserendpoints(t *testing.T) {
	t.Run("Get Username: Pass", testPassGetUsername)
	t.Run("Get Username: no Username", testFailGetUsername)
	t.Run("Update Username: Pass", testPassUpdateUser)
	t.Run("Update Username:Fail", testFailUpdateUsername)
	t.Run("Delete Username: Pass", testPassDeleteUser)
	t.Run("Delete Username:Fail", testFailDeleteUsername)
}

func testPassGetUsername(t *testing.T) {
	c := Connection{
		db: &dbPassMock{},
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/getUser/captainnobody1@email.com", nil)
	c.getUser(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func testFailGetUsername(t *testing.T) {
	c := Connection{
		db: &dbFailMock{},
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/getUser/captainnobody1@email.com", nil)
	c.getUser(w, req)
	resp := w.Result()
	if resp.StatusCode != 500 {
		t.Fatalf("Expected status code 500, got %d", resp.StatusCode)
	}
}

func testPassUpdateUser(t *testing.T) {
	c := Connection{
		db: &dbPassMock{},
	}
	w := httptest.NewRecorder()
	err := json.NewEncoder(w).Encode(User{
		ID:    1,
		Name:  "Miriah Peterson",
		Email: "captainnboody1@gmail.com",
	})
	if err != nil {
		t.Fatalf("Error encoding JSON")
	}
	req := httptest.NewRequest("POST", "/updateUser/captainnobody1@email.com", nil)
	c.updateUser(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		fmt.Println(resp)
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func testFailUpdateUsername(t *testing.T) {
	c := Connection{
		db: &dbFailMock{},
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/updateUser/captainnobody1@email.com", nil)
	c.updateUser(w, req)
	resp := w.Result()
	if resp.StatusCode != 500 {
		t.Fatalf("Expected status code 500, got %d", resp.StatusCode)
	}
}

func testPassDeleteUser(t *testing.T) {
	c := Connection{
		db: &dbPassMock{},
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/deleteUser/captainnobody1@gmail.com", nil)
	c.deleteUser(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func testFailDeleteUsername(t *testing.T) {
	c := Connection{
		db: &dbFailMock{},
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/deleteUser/captainnobody1@gmail.com", nil)
	c.deleteUser(w, req)
	resp := w.Result()
	if resp.StatusCode != 500 {
		t.Fatalf("Expected status code 500, got %d", resp.StatusCode)
	}
}
