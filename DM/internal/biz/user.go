package biz

import (
	"DM/internal/models"
	"DM/internal/token"
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserRepository interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, phone string) (models.User, error)
	//UpdateToken(ctx context.Context, refresh_token string) error
	FindAll(ctx context.Context) ([]models.User, error)
	Delete(ctx context.Context, id uint) error
}

type UserService struct {
	userRepo     UserRepository
	employeeRepo EmployeeRepo
}

func NewUserService(userRepo UserRepository, employeeRepo EmployeeRepo) *UserService {
	return &UserService{userRepo, employeeRepo}
}

func (s *UserService) SignUpUser(ctx context.Context, phone, password, confirmpass string, isemployee bool, role string) (uint32, string, string, bool, string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	password = string(hash)

	var employeeID = randomIDEm()
	if isemployee {
		employee := &models.Employee{
			//IDEm:            uint32(employeeID),
			Name:            "van",
			Phone:           phone,
			Password:        password,
			Role:            role,
			Salary:          0,
			SubDepartmentID: 2,
		}
		if err := s.employeeRepo.Create(ctx, employee); err != nil {
			return 0, "", "", false, ""
		}
	}
	user := &models.User{Phone: phone, Password: password, ConfirmPass: password, IsEmployee: isemployee, Role: role}
	user.IDUser = uint32(employeeID)
	if err := s.userRepo.SignUp(ctx, user); err != nil {
		return 0, "", "", false, ""
	}
	return user.IDUser, user.Phone, user.Password, user.IsEmployee, user.Role
}

func (s *UserService) SignInUser(ctx context.Context, phone, password, role string) (string, string, error) {
	user, err := s.userRepo.SignIn(ctx, phone)
	if err != nil {
		return "", "", errors.New("invalid phone or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid phone or password compare")
	}

	if user.IsEmployee {
		emp, err := s.employeeRepo.FindByPhone(ctx, phone)
		if err != nil {
			return "", "", errors.New("invalid phone or password employee")
		}
		err = bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(password))
		if err != nil {
			return "", "", errors.New("invalid phone or password compare")
		}
	}

	accessToken, err := token.GenerateAccessToken(user.Phone, role, time.Minute*15)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := token.GenerateRefreshToken(user.Phone, time.Hour*24)
	if err != nil {
		return "", "", err
	}
	// user.AccessToken = accessToken
	// user.RefreshToken = refreshToken
	return accessToken, refreshToken, nil
}

func (s *UserService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {

	claims, err := token.ParseJWT(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	accessToken, err := token.GenerateAccessToken(claims.Phone, claims.Role, time.Minute*15)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *UserService) GetAllUser(ctx context.Context) ([]models.User, error) {
	return s.userRepo.FindAll(ctx)
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}

func randomIDEm() uint {
	return uint(rand.Intn(1000-1) + 1)
}
