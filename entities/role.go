package entities

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type RoleType string

const (
	ADMIN   RoleType = "admin"
	MAKER   RoleType = "maker"
	CHECKER RoleType = "checker"
	SIGNER  RoleType = "signer"
	VIEWER  RoleType = "viewer"
)

type Role struct {
	//ID        string         `gorm:"type:binary(16);primaryKey" json:"id"`
	ID        int            `json:"id"`
	Title     RoleType       `gorm:"type:enum('admin', 'maker', 'checker', 'signer', 'viewer');default:'viewer'" json:"title"`
	Active    bool           `json:"-"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (roleType *RoleType) Scan(value interface{}) error {
	*roleType = RoleType(value.([]byte))
	return nil
}

func (roleType RoleType) Value() (driver.Value, error) {
	return string(roleType), nil
}

// func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
// 	tx.Statement.SetColumn("ID", uuid.NewV4())
// 	return
// }
