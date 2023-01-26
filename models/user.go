// Create package models, import "time",
package models

import "time"

// User struct, UsersProfileResponse  struct,
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// and UsersProfileResponse TableName method here ...
func (UsersProfileResponse) TableName() string {
	return "users"
}
