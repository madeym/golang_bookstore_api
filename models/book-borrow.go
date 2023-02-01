package models

import (
	"belajar-api/helper"
	"time"

	"gorm.io/gorm"
)

type BookBorrowScan struct {
	ID               int
	UserID           int
	User             User `gorm:"foreignKey:UserID"`
	BookID           int
	Book             Book `gorm:"embedded;embeddedPrefix:book_;foreignKey:BookID"`
	BorrowTime       time.Time
	ExpireBorrowTime time.Time
}

type BookBorrow struct {
	ID               int
	UserID           int
	BookID           int
	BorrowTime       time.Time
	ExpireBorrowTime time.Time
}

func init() {
	db.AutoMigrate(&BookBorrow{})
}

func GetDefaultQueryBookBorrow() *gorm.DB {
	return db.
		Table("book_borrows").
		Select("book_borrows.id, book_borrows.user_id, book_borrows.book_id, book_borrows.borrow_time, book_borrows.expire_borrow_time, User.id AS User__id, User.name AS User__name, User.email AS User__email, Book.id AS Book__id, Book.title AS book_title, Book.description AS book_description, Book.price AS book_price, Book.rating AS book_rating, Book.discount AS book_discount, Book.quantity AS book_quantity, Book.genre_id AS book_genre_id, Genre.id AS Genre__id, Genre.name AS Genre__name, Genre.description AS Genre__description, Book.status_id AS book_status_id, Status.id AS Status__id, Status.name AS Status__name").
		Joins("LEFT JOIN users User ON book_borrows.user_id = User.id LEFT JOIN books Book ON book_borrows.book_id = Book.id LEFT JOIN book_statuses Status ON Book.status_id = Status.id LEFT JOIN genres Genre ON Book.genre_id = Genre.id")
}

func GetBookBorrows(keywordName []string, keyword string, param []string, orderName string, orderBy string, pagination helper.Pagination) ([]BookBorrowScan, error) {
	var b []BookBorrowScan
	mainQuery := GetDefaultQueryBookBorrow()
	keywordQuery := helper.AddKeywordQuery(db, keywordName, keyword)
	paramQuery := helper.AddParamQuery(db, param)
	mainQuery = helper.AddOrderQuery(mainQuery, orderName, orderBy)
	mainQuery = helper.AddPaginationQuery(mainQuery, pagination)

	err := mainQuery.Where(keywordQuery).Where(paramQuery).Scan(&b).Error
	return b, err
}

func CreateBorrowBook(ID int) (BookBorrowScan, string, error) {
	var b Book
	var bbc BookBorrow
	var bbs BookBorrowScan
	err := db.First(&b, ID).Error
	if err == nil {
		if b.StatusID == 3 || b.StatusID == 2 {
			return bbs, "Book Not Available", err
		} else {
			b.StatusID = 3
			err = db.Model(&b).Where("id=?", ID).Updates(b).Error
			if err == nil {
				bbc.UserID = helper.UserData.ID
				bbc.BookID = ID
				bbc.BorrowTime = time.Now()
				bbc.ExpireBorrowTime = bbc.BorrowTime.Add(time.Hour)

				err = db.Create(&bbc).Error
				if err == nil {
					err = GetDefaultQueryBookBorrow().First(&bbs, bbc.ID).Error
					return bbs, "", err
				}
				return bbs, "", err
			}
			return bbs, "", err
		}
	}
	return bbs, "", err
}

func ReturnBookBorrow(ID int) (BookBorrowScan, string, error) {
	var b Book
	var bb BookBorrow
	var bbs BookBorrowScan

	err := db.Where("user_id = ? AND book_id = ?", helper.UserData.ID, ID).First(&bb).Error
	if err == nil {
		err = db.Model(&b).Where("id = ?", bb.BookID).Updates(Book{StatusID: 1}).Error
		if err == nil {
			err = GetDefaultQueryBookBorrow().Where("book_borrows.book_id = ? AND book_borrows.id = ?", bb.BookID, bb.ID).Scan(&bbs).Error
			if err == nil {
				db.Delete(&bb, bb.ID)
				if err != nil {
					err = db.Model(&b).Where("id = ?", bb.BookID).Updates(Book{StatusID: 3}).Error
				}
				return bbs, "", err
			}
			return bbs, "", err
		}
		return bbs, "", err
	}
	return bbs, "", err
}
