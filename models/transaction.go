package models

import (
	"belajar-api/helper"
	"time"

	"gorm.io/gorm"
)

type TransactionScan struct {
	ID                 int
	UserID             int
	User               User `gorm:"ForeignKey:UserID; references:ID"`
	BookID             int
	Book               Book `gorm:"embedded;embeddedPrefix:book_;ForeignKey:BookID; references:ID"`
	Quantity           int
	PaymentBankID      int
	PaymentBank        PaymentBank `gorm:"ForeignKey:PaymentBankID; references:ID"`
	PaymentBankAccount string
	PaymentDate        time.Time
}

type Transaction struct {
	ID                 int
	UserID             int
	BookID             int
	Quantity           int
	PaymentBankID      int
	PaymentBankAccount string
	PaymentDate        time.Time
}

func init() {
	db.AutoMigrate(&Transaction{})
}

func GetDefaultQueryTransaction() *gorm.DB {
	var defaultQueary = db.
		Table("transactions").
		Select("transactions.id, transactions.user_id, User.id AS User__id, User.name AS User__name, User.email AS User__email, transactions.book_id, Book.id AS Book__id, Book.title AS book_title, Book.description AS book_description, Book.price AS book_price, Book.rating AS book_rating, Book.discount AS book_discount, Book.quantity AS book_quantity, Book.genre_id AS book_genre_id, Genre.id AS Genre__id, Genre.name AS Genre__name, Genre.description AS Genre__description, transactions.quantity, Book.status_id AS book_status_id, Status.id AS Status__id, Status.name AS Status__name, transactions.payment_bank_id, Bank.id AS PaymentBank__id, Bank.name AS PaymentBank__name, transactions.payment_bank_account, transactions.payment_date").
		Joins("LEFT JOIN users User ON transactions.user_id = User.id LEFT JOIN books Book ON transactions.book_id = Book.id LEFT JOIN book_statuses Status ON Book.status_id = Status.id LEFT JOIN genres Genre ON Book.genre_id = Genre.id LEFT JOIN banks Bank ON transactions.payment_bank_id = Bank.id")

	return defaultQueary
}

func GetAllTransactions(pagination helper.Pagination, param []string, keywordName []string, keyword string) ([]TransactionScan, error) {
	var t []TransactionScan
	var err error
	mainQuery := GetDefaultQueryTransaction()
	keywordQuery := helper.AddKeywordQuery(mainQuery, keywordName, keyword)
	paramQuery := helper.AddParamQuery(mainQuery, param)
	mainQuery = helper.AddPaginationQuery(mainQuery, pagination)

	err = mainQuery.Where(keywordQuery).Where(paramQuery).Scan(&t).Error

	return t, err
}

func GetTransactionByID(ID int) (TransactionScan, error) {
	var t Transaction
	var ts TransactionScan

	err := GetDefaultQueryTransaction().First(&t, ID).Scan(&ts).Error

	return ts, err
}
