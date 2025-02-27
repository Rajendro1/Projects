package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Address struct {
	Village    Village `json:"village"`
	PostOffice string  `json:"post_office"`
}
type Village struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}
type User struct {
	ID      uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string  `json:"name" gorm:"not null"`
	Email   string  `json:"email" gorm:"unique;not null"`
	Address Address `json:"address" gorm:"type:json"`
}

// Implement GORM serialization using generic functions
func (a Address) Value() (driver.Value, error) {
	return ToJSON(a)
}

func (a *Address) Scan(value interface{}) error {
	return FromJSON(value, a)
}

func ToJSON(value interface{}) (driver.Value, error) {
	return json.Marshal(value) // Convert struct to JSON before saving
}

// Generic function to implement Scan() for GORM deserialization
func FromJSON(value interface{}, target interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan JSON data")
	}
	return json.Unmarshal(bytes, target) // Convert JSON back to struct
}

// TableName overrides GORM's default naming convention
func (User) TableName() string {
	return "users"
}
