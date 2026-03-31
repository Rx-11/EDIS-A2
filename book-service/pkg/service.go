package pkg

import (
	"github.com/Rx-11/EDIS-A2/book-service/pkg/books"
)

var (
	BookRepo books.BookRepo
)

func init() {
	BookRepo = books.NewBookMySQLRepo()
}
