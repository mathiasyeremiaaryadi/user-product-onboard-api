package entities

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// RoleID         string         `gorm:"type:binary(16)" json:"role_id"`
	ID             int            `json:"id"`
	PersonalNumber string         `gorm:"type:varchar(20)" json:"personalNumber"`
	Password       string         `gorm:"type:varchar(255)" json:"password"`
	Email          string         `gorm:"type:varchar(50)" json:"email"`
	Name           string         `gorm:"type:varchar(50)" json:"name"`
	RoleID         int            `json:"roleID"`
	Role           Role           `json:"role"`
	Active         bool           `gorm:"default:false" json:"active"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"upatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}

type UserList struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RoleID int    `json:"-"`
	Role   Role   `json:"role"`
	Active bool   `json:"active"`
}

type UserDetail struct {
	ID             int    `json:"id"`
	PersonalNumber string `json:"personalNumber"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	RoleID         int    `json:"-"`
	Role           Role   `json:"role"`
	Active         bool   `json:"active"`
}

type UserProduct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword)
}

func (UserProduct) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// tx.Statement.SetColumn("ID", uuid.NewV4())
	user.Password = HashPassword(user.Password)
	user.Name = strings.ToLower(user.Name)

	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if user.Password == "" {
		return
	}

	user.Password = HashPassword(user.Password)
	user.Name = strings.ToLower(user.Name)

	if !user.Active {
		user.Active = false
	} else {
		user.Active = true
	}

	return
}
