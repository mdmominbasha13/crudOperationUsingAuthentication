package models

import (
	"github.com/CrudOperationUsingAuthentication/pkg/database"
	"github.com/dgriJalva/jwt-go"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

type Book struct {
	gorm.Model
	Name        string `gorm:"json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate()
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books

}
func GetBookById(id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("ID=?", id).Find(&getbook)
	return &getbook, db
}

func DeleteBookById(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book

}
