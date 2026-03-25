package pkg

import (
	"github.com/Rx-11/EDIS-A2/book-web-bff/pkg/books"
	"github.com/Rx-11/EDIS-A2/book-web-bff/pkg/users"
)

var (
	UserRepo users.UserRepo
	BookRepo books.BookRepo
)

func init() {
	UserRepo = users.NewUserMySQLRepo()
	BookRepo = books.NewBookMySQLRepo()
}
