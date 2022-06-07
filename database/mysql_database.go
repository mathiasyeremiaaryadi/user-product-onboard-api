package database

import (
	"fmt"
	"log"
	"user-product-service/config"
	"user-product-service/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDatabase struct {
	db  *gorm.DB
	err error
}

func (mysqldb *MySQLDatabase) SetDBConnection() {
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.ENV["DB_USERNAME"], config.ENV["DB_HOST"], config.ENV["DB_PORT"], config.ENV["DB_NAME"])

	mysqldb.db, mysqldb.err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if mysqldb.err != nil {
		panic("Failed to connect database")
	} else {
		log.Println("Database is connected")
	}

	mysqldb.db.Migrator().DropTable(&entities.Role{}, &entities.User{}, &entities.Product{})
	mysqldb.db.AutoMigrate(&entities.Role{}, &entities.User{}, &entities.Product{})
}

func (mysqldb *MySQLDatabase) SetDBSeed() {
	mysqldb.db.Create(&[]entities.Role{
		{
			Title:  "admin",
			Active: true,
		},
		{
			Title:  "maker",
			Active: true,
		},
		{
			Title:  "checker",
			Active: true,
		},
		{
			Title:  "signer",
			Active: true,
		},
		{
			Title:  "viewer",
			Active: true,
		},
	})

	mysqldb.db.Create(&[]entities.User{
		{
			PersonalNumber: "123",
			Password:       "secret123",
			Email:          "user.a@gmail.com",
			Name:           "User A",
			RoleID:         1,
			Active:         true,
		},
		{
			PersonalNumber: "345",
			Password:       "secret123",
			Email:          "user.b@gmail.com",
			Name:           "User B",
			RoleID:         2,
			Active:         true,
		},
		{
			PersonalNumber: "567",
			Password:       "secret123",
			Email:          "user.c@gmail.com",
			Name:           "User C",
			RoleID:         2,
			Active:         false,
		},
		{
			PersonalNumber: "789",
			Password:       "secret123",
			Email:          "user.d@gmail.com",
			Name:           "User D",
			RoleID:         3,
			Active:         true,
		},
		{
			PersonalNumber: "91011",
			Password:       "secret123",
			Email:          "user.E@gmail.com",
			Name:           "User E",
			RoleID:         3,
			Active:         false,
		},
		{
			PersonalNumber: "111213",
			Password:       "secret123",
			Email:          "user.f@gmail.com",
			Name:           "User F",
			RoleID:         4,
			Active:         true,
		},
		{
			PersonalNumber: "131415",
			Password:       "secret123",
			Email:          "user.g@gmail.com",
			Name:           "User G",
			RoleID:         4,
			Active:         false,
		},
		{
			PersonalNumber: "151617",
			Password:       "secret123",
			Email:          "user.h@gmail.com",
			Name:           "User H",
			RoleID:         5,
			Active:         true,
		},
		{
			PersonalNumber: "171819",
			Password:       "secret123",
			Email:          "user.i@gmail.com",
			Name:           "User I",
			RoleID:         5,
			Active:         false,
		},
	})

	mysqldb.db.Create(&[]entities.Product{
		{
			Name:        "Product A",
			Description: "A very good and cheap product",
			Status:      "inactive",
			MakerID:     2,
			CheckerID:   3,
			SignerID:    4,
		},
	})
}

func (mysqldb MySQLDatabase) GetDBConnection() *gorm.DB {
	return mysqldb.db
}
