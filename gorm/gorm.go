package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string
	Books []*Book `gorm:"many2many:authors_books;"`
}

type Book struct {
	gorm.Model
	Title   string
	Authors []*Author `gorm:"many2many:authors_books;"`
}

func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
	return db, nil
}
