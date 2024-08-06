package data

import (
	"DM/internal/biz"
	"DM/internal/entity"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *Data) biz.UserRepository {
	return &userRepository{db: db.DB}
}

func (r *userRepository) SignUp(ctx context.Context, user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) SignIn(ctx context.Context, phone string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	return user, err
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *userRepository) GetByPage(ctx context.Context, page, limit uint32) ([]*entity.User, error) {
	var users []*entity.User
	if int(page) <= 0 {
		page = 1
	}
	err := r.db.Offset(int(limit * (page - 1))).Limit(int(limit)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
