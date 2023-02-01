package models

import (
	"belajar-api/helper"
	"belajar-api/request"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Book struct {
	ID          int
	CoverImage  string
	Title       string
	Author      string
	PublishDate time.Time
	Page        int
	Description string
	Price       int
	Rating      int
	Discount    int
	Quantity    int
	GenreID     int
	StatusID    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Genre       Genre      `gorm:"foreignKey:GenreID"`
	Status      BookStatus `gorm:"foreignKey:StatusID"`
}

var db *gorm.DB

func init() {
	db.AutoMigrate(&Book{})
}

func GetDefaultQueryBook() *gorm.DB {
	return db.Joins("Genre").Joins("Status")
}

func GetAllBooks(keywordName []string, keyword string, param []string, orderName string, orderBy string, pagination helper.Pagination) ([]Book, error) {
	var books []Book
	mainQuery := GetDefaultQueryBook()
	keywordQuery := helper.AddKeywordQuery(db, keywordName, keyword)
	paramQuery := helper.AddParamQuery(db, param)
	mainQuery = helper.AddOrderQuery(mainQuery, orderName, orderBy)
	mainQuery = helper.AddPaginationQuery(mainQuery, pagination)

	err := mainQuery.Where(keywordQuery).Where(paramQuery).Find(&books).Error

	return books, err
}

func FindBookByID(ID int) (Book, error) {
	var book Book
	err := GetDefaultQueryBook().First(&book, ID).Error

	return book, err
}

func CreateBook(book Book) (Book, error) {
	var g Genre
	var s BookStatus
	err := db.First(&g, book.GenreID).Error
	if err == nil {
		err := db.First(&s, book.StatusID).Error
		if err == nil {
			err := db.Create(&book).Error
			if err == nil {
				err = GetDefaultQueryBook().First(&book, book.ID).Error
			}
			return book, err
		}
		return book, err
	}
	return book, err
}

func UpdateBook(ID int, book Book) (Book, error) {
	var g Genre
	var s BookStatus
	err := db.First(&g, book.GenreID).Error
	if err == nil {
		err := db.First(&s, book.StatusID).Error
		if err == nil {
			err := db.Model(&book).Where("id=?", ID).Updates(&book).Error
			if err == nil {
				err = GetDefaultQueryBook().First(&book, ID).Error
			}
			return book, err
		}
		return book, err
	}
	return book, err
}

func DeleteBook(ID int) error {
	var book Book
	err := db.Delete(&book, ID).Error

	return err
}

func BuyBook(ID int, request *request.BuyBookRequest) (TransactionScan, Book, string, error) {
	var book Book
	var t Transaction
	var ts TransactionScan

	err := db.First(&book, ID).Error
	if err == nil {
		if book.Quantity == 0 {
			return ts, book, "Book Sold Out", nil
		} else {
			book.Quantity = book.Quantity - 1
			user := helper.UserData
			t.UserID = user.ID
			t.BookID = book.ID
			t.Quantity = request.Quantity
			t.PaymentBankID = request.PaymentBankID
			t.PaymentBankAccount = request.PaymentBankAccount
			t.PaymentDate = time.Now()

			err = db.Omit(clause.Associations).Create(&t).Error
			if err == nil {
				err = db.Model(&book).Select("Quantity").Updates(Book{Quantity: book.Quantity}).Error
				if err == nil {
					err = GetDefaultQueryTransaction().Where("transactions.id = ?", t.ID).Scan(&ts).Error
					return ts, book, "", err
				}
				_ = db.Delete(&t, t.ID).Error
				return ts, book, "", err
			}
			return ts, book, "", err
		}
	}
	return ts, book, "", err
}
