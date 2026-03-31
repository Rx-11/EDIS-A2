package users

import (
	"github.com/Rx-11/EDIS-A2/customer-service/pkg/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	FetchUserByID(db *gorm.DB, id uint) (*models.User, error)
	FetchUserByUserID(db *gorm.DB, userid string) (*models.User, error)
	CreateUser(db *gorm.DB, user models.User) (*models.User, error)
}
