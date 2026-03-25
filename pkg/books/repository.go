package books

import (
	"github.com/Rx-11/EDIS-A1/pkg/models"
	"gorm.io/gorm"
)

type BookRepo interface {
	FetchBookByISBN(db *gorm.DB, isbn string) (*models.Book, error)
	CreateBook(db *gorm.DB, book models.Book) (*models.Book, error)
	UpdateBook(db *gorm.DB, book models.Book) (*models.Book, error)
}
