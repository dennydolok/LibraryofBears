package Models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Role      uint      `json:"role"`
	Password  string    `json:"password,omitempty" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
