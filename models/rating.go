package models

import (
	"belajar-api/helper"
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	ID        int
	BookID    int
	UserID    int
	Rating    int
	Review    string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"ForeignKey:UserID;references:ID"`
}

func init() {
	db.AutoMigrate(&Rating{})
}

func GetDefaultQueryRating() *gorm.DB {
	return db.Joins("User")
}

func GetAllRatings(book_id int, keywordName []string, keyword string, orderName string, orderBy string, pagination helper.Pagination) ([]Rating, error) {
	var ratings []Rating
	mainQuery := GetDefaultQueryRating()
	keywordQuery := helper.AddKeywordQuery(db, keywordName, keyword)
	mainQuery = helper.AddOrderQuery(mainQuery, orderName, orderBy)
	mainQuery = helper.AddPaginationQuery(mainQuery, pagination)

	err := mainQuery.Where("book_id = ?", book_id).Where(keywordQuery).Find(&ratings).Error

	return ratings, err
}

func CreateRating(rating Rating) (Rating, error) {
	var b Book
	err := db.First(&b, rating.BookID).Error
	if err == nil {
		rating.UserID = helper.UserData.ID
		err = db.Create(&rating).Error
		if err == nil {
			err = db.First(&rating, rating.ID).Error
		}
	}

	return rating, err
}
