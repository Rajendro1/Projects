package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
}

// TableName overrides GORM's default naming convention
func (User) TableName() string {
	return "users"
}
