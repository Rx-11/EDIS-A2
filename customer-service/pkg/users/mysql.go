package users

import (
	"github.com/Rx-11/EDIS-A1/pkg/models"
	"gorm.io/gorm"
)

type userRepo struct {
}

func (r *userRepo) CreateUser(db *gorm.DB, user models.User) (*models.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FetchUserByID(db *gorm.DB, id uint) (*models.User, error) {
	user := &models.User{ID: id}
	err := db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) FetchUserByUserID(db *gorm.DB, userid string) (*models.User, error) {
	user := &models.User{UserID: userid}
	err := db.Where("user_id = ?", userid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserMySQLRepo() UserRepo {
	return &userRepo{}
}
