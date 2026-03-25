package books

import (
	"github.com/Rx-11/EDIS-A2/book-service/pkg/models"
	"gorm.io/gorm"
)

type bookRepo struct {
}

func (r *bookRepo) CreateBook(db *gorm.DB, book models.Book) (*models.Book, error) {
	err := db.Create(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepo) FetchBookByISBN(db *gorm.DB, isbn string) (*models.Book, error) {
	book := &models.Book{ISBN: isbn}
	err := db.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepo) UpdateBook(db *gorm.DB, book models.Book) (*models.Book, error) {
	err := db.Save(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func NewBookMySQLRepo() BookRepo {
	return &bookRepo{}
}
