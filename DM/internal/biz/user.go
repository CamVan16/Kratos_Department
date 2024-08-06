package biz

import (
	"DM/internal/entity"
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserRepository interface {
	SignUp(ctx context.Context, user *entity.User) error
	SignIn(ctx context.Context, phone string) (entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	Delete(ctx context.Context, id uint) error
	GetByPage(ctx context.Context, page, limit uint32) ([]*entity.User, error)
}

type UserUC struct {
	userRepo     UserRepository
	employeeRepo EmployeeRepo
}

func NewUserUC(userRepo UserRepository, employeeRepo EmployeeRepo) *UserUC {
	return &UserUC{userRepo, employeeRepo}
}

func (uc *UserUC) SignUpUser(ctx context.Context, phone, password, confirmpass string, isemployee bool, role string) (uint32, string, string, bool, string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	password = string(hash)

	var employeeID = randomIDEm()
	if isemployee {
		employee := &entity.Employee{
			//IDEm:            uint32(employeeID),
			Name:            "van",
			Phone:           phone,
			Password:        password,
			Role:            role,
			Salary:          0,
			SubDepartmentID: 2,
		}
		if err := uc.employeeRepo.Create(ctx, employee); err != nil {
			return 0, "", "", false, ""
		}
	}
	user := &entity.User{Phone: phone, Password: password, ConfirmPass: password, IsEmployee: isemployee, Role: role}
	user.IDUser = uint32(employeeID)
	if err := uc.userRepo.SignUp(ctx, user); err != nil {
		return 0, "", "", false, ""
	}
	return user.IDUser, user.Phone, user.Password, user.IsEmployee, user.Role
}

func (uc *UserUC) SignInUser(ctx context.Context, phone, password, role string) (string, string, error) {
	user, err := uc.userRepo.SignIn(ctx, phone)
	if err != nil {
		return "", "", errors.New("invalid phone or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid phone or password compare")
	}

	if user.IsEmployee {
		emp, err := uc.employeeRepo.FindByPhone(ctx, phone)
		if err != nil {
			return "", "", errors.New("invalid phone or password employee")
		}
		err = bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(password))
		if err != nil {
			return "", "", errors.New("invalid phone or password compare")
		}
	}

	accessToken, err := GenerateAccessToken(user.Phone, role, time.Minute*15)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(user.Phone, time.Hour*24)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (uc *UserUC) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {

	claims, err := ParseJWT(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	accessToken, err := GenerateAccessToken(claims.Phone, claims.Role, time.Minute*15)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (uc *UserUC) GetAllUser(ctx context.Context) ([]entity.User, error) {
	return uc.userRepo.FindAll(ctx)
}

func (uc *UserUC) DeleteUser(ctx context.Context, id uint) error {
	return uc.userRepo.Delete(ctx, id)
}

func (uc *UserUC) GetUserByPage(ctx context.Context, page, limit uint32) ([]*entity.User, error) {
	return uc.userRepo.GetByPage(ctx, page, limit)
}

func randomIDEm() uint {
	return uint(rand.Intn(1000-1) + 1)
}
