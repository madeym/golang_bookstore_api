package models

import (
	"belajar-api/helper"
	"time"
)

type Genre struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func init() {
	db.AutoMigrate(&Book{})
}

func GetAllGenres(keywordName []string, keyword string, param []string, orderName string, orderBy string, pagination helper.Pagination) ([]Genre, error) {
	var genre []Genre
	mainQuery := db
	keywordQuery := helper.AddKeywordQuery(db, keywordName, keyword)
	paramQuery := helper.AddParamQuery(db, param)
	mainQuery = helper.AddOrderQuery(mainQuery, orderName, orderBy)
	mainQuery = helper.AddPaginationQuery(mainQuery, pagination)

	err := mainQuery.Where(keywordQuery).Where(paramQuery).Find(&genre).Error

	return genre, err
}

func GetGenreByID(ID int) (Genre, error) {
	var genre Genre
	err := db.Find(&genre, ID).Error

	return genre, err
}

func CreateGenre(newGenre Genre) (Genre, error) {
	err := db.Create(&newGenre).Error
	if err == nil {
		err = db.First(&newGenre, newGenre.ID).Error
		return newGenre, err
	}

	return newGenre, err
}

func UpdateGenre(updateGenre Genre) (Genre, error) {
	err := db.Model(&updateGenre).Where("id = ?", updateGenre.ID).Updates(&updateGenre).Error

	if err == nil {
		err = db.First(&updateGenre, updateGenre.ID).Error
		return updateGenre, err
	}

	return updateGenre, err
}

func DeleteGenre(id int) error {
	var g Genre
	err := db.Debug().Delete(&g, id).Error

	return err
}
