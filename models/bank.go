package models

import "belajar-api/config"

type PaymentBank struct {
	ID   int
	Name string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&PaymentBank{})
}
