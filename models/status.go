package models

type BookStatus struct {
	ID   int
	Name string
}

func init() {
	db.AutoMigrate(&BookStatus{})
}

func GetAllBookStatus() ([]BookStatus, error) {
	var bookStatus []BookStatus
	err := db.Find(&bookStatus).Error

	return bookStatus, err
}
