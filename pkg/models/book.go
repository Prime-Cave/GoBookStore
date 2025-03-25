package models

import (
	"github.com/jinzhu/gorm"
	"go/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)// This will first check if the record exists, if it does it will give an error
	db.Create(&b) // This will insert a new record into the database
	return b
}

func GetAllBooks() []Book{
	var Books []Book // This an array of Book
	db.Find(&Books) // This will return all the books by querying the database for all the books
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
