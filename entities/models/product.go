package models

import (
	"database/sql/driver"
	"strings"
	"time"

	"gorm.io/gorm"
)

type StatusType string

const (
	INACTIVE StatusType = "inactive"
	APPROVED StatusType = "approved"
	active   StatusType = "active"
)

type Product struct {
	ID          int            `json:"id"`
	Name        string         `gorm:"type:varchar(50)" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Status      StatusType     `gorm:"type:enum('inactive', 'approved', 'active');default:'inactive'" json:"status"`
	MakerID     int            `json:"maker_id"`
	CheckerID   int            `json:"checker_id"`
	SignerID    int            `json:"signer_id"`
	Maker       User           `json:"maker"`
	Checker     User           `json:"checker"`
	Signer      User           `json:"signer"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type ProductList struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      StatusType `json:"status"`
}

type ProductDetail struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Status      StatusType  `json:"status"`
	MakerID     int         `json:"-"`
	CheckerID   int         `json:"-"`
	SignerID    int         `json:"-"`
	Maker       UserProduct `json:"maker"`
	Checker     UserProduct `json:"checker"`
	Signer      UserProduct `json:"signer"`
}

func (statusType *StatusType) Scan(value interface{}) error {
	*statusType = StatusType(value.([]byte))
	return nil
}

func (statusType StatusType) Value() (driver.Value, error) {
	return string(statusType), nil
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	//tx.Statement.SetColumn("ID", uuid.NewV4())
	product.Name = strings.ToLower(product.Name)
	return
}

func (product *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	//tx.Statement.SetColumn("ID", uuid.NewV4())
	product.Name = strings.ToLower(product.Name)
	return
}
