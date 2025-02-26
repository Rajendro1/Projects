package users

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
	"main.go/models"
)

type UserService struct {
	DB *gorm.DB
}

// CreateUser method for UserService
func (s *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := s.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
