package pkg

import (
	"github.com/Rx-11/EDIS-A2/customer-service/pkg/users"
)

var (
	UserRepo users.UserRepo
)

func init() {
	UserRepo = users.NewUserMySQLRepo()
}
