package users

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
)

// JSONMatcher ensures JSON arguments match when stored in DB.
type JSONMatcher struct {
	Value interface{}
}

func (jm JSONMatcher) Match(v driver.Value) bool {
	bytes, ok := v.([]byte)
	if !ok {
		return false
	}

	var actual, expected interface{}
	if err := json.Unmarshal(bytes, &actual); err != nil {
		return false
	}

	expectedBytes, _ := json.Marshal(jm.Value)
	if err := json.Unmarshal(expectedBytes, &expected); err != nil {
		return false
	}

	// Convert both values to JSON string and compare
	return string(bytes) == string(expectedBytes)
}

func TestCreateUser(t *testing.T) {
	// Create mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock DB: %v", err)
	}
	defer mockDB.Close()

	// Initialize GORM with mock DB
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error initializing GORM with mock DB: %v", err)
	}

	// Define sample user with Address (nested JSON)
	address := models.Address{
		Village:    models.Village{OldName: "Old Village", NewName: "New Village"},
		PostOffice: "Kolkata Post",
	}

	// Expect DB insert with JSON column
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs("John Doe", "john@example.com", JSONMatcher{Value: address}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Create HTTP request
	reqBody := `{
		"name": "John Doe",
		"email": "john@example.com",
		"address": {
			"village": {
				"old_name": "Old Village",
				"new_name": "New Village"
			},
			"post_office": "Kolkata Post"
		}
	}`
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder (simulates HTTP response)
	w := httptest.NewRecorder()

	// Call the CreateUser function (assumed to exist)
	userService := &UserService{DB: gormDB}
	userService.CreateUser(w, req)

	// Validate HTTP response
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Ensure all SQL expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("SQL expectations were not met: %v", err)
	}
}
