package users

import (
	"net/http"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Test cases using sqlmock
func TestCreateUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock DB: %v", err)
	}
	defer mockDB.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error initializing GORM with mock DB: %v", err)
	}
	gormDB = gormDB.Debug()

	userService := &UserService{DB: gormDB}

	mock.ExpectBegin() // GORM uses transactions internally
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs("John Doe", "john@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	reqBody := `{"name": "John Doe", "email": "john@example.com"}`
	req, err := http.NewRequest("POST", "/users", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	w := &mockResponseWriter{headers: http.Header{}}
	// Create UserService instance with mock DB
	userService.CreateUser(w, req)

	if w.statusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.statusCode)
	}
	// Ensure expectations are met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("SQL expectations were not met: %v", err)
	}
}

type mockResponseWriter struct {
	headers    http.Header
	body       []byte
	statusCode int
}

func (m *mockResponseWriter) Header() http.Header {
	return m.headers
}

func (m *mockResponseWriter) Write(b []byte) (int, error) {
	m.body = append(m.body, b...)
	return len(b), nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}
