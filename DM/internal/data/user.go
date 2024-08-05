package data

import (
	"DM/internal/biz"
	"DM/internal/models"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *Data) biz.UserRepository {
	return &userRepository{db: db.DB}
}

func (r *userRepository) SignUp(ctx context.Context, user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) SignIn(ctx context.Context, phone string) (models.User, error) {
	var user models.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	return user, err
}

func (r *userRepository) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// func (r *userRepository) UpdateToken(ctx context.Context, refresh_token string) (string, error){

// }
