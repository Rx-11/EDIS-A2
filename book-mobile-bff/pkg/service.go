package pkg

import (
	"github.com/Rx-11/EDIS-A1/pkg/books"
	"github.com/Rx-11/EDIS-A1/pkg/users"
)

var (
	UserRepo users.UserRepo
	BookRepo books.BookRepo
)

func init() {
	UserRepo = users.NewUserMySQLRepo()
	BookRepo = books.NewBookMySQLRepo()
}
