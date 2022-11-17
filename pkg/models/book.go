package models

import (
	"github.com/rahul/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model // model is used to strucutre the data  , and also will see its use after running the app
	// by deleting it
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect() // a func. which helps us to connect to database
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

// the user first interacts with routes and then routes
//send control to controllers where we have build all the logic
//and then controller have to perform some operations with database
//the operations have to reside in models file which is book.go in this case
//means we have to have a different function for a different controllers that we create

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // newrecord is a func of gorm which is used to interact with database
	db.Create(&b)
	return b
}
func GetAllBooks() []Book { // might be wrong coz the name of routes is getbook and here he has used getallbooks
	var Books []Book // slice to store all books
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
