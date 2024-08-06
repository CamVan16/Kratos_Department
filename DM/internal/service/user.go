package service

import (
	v1 "DM/api/user/v1"
	"DM/internal/entity"
	"context"
	"errors"
)

type userUC interface {
	SignUpUser(ctx context.Context, phone, password, confirmpass string, isemployee bool, role string) (uint32, string, string, bool, string)
	SignInUser(ctx context.Context, phone, password, role string) (string, string, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, id uint) error
	GetUserByPage(ctx context.Context, page, limit uint32) ([]*entity.User, error)
}
type UserService struct {
	v1.UnimplementedUserServiceServer
	//uc *biz.UserUC
	uc userUC
}

func NewUserService(uc userUC) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) SignUpUser(ctx context.Context, req *v1.SignUpUserRequest) (*v1.SignUpUserRespone, error) {
	id, phone, password, isemployee, role := s.uc.SignUpUser(ctx, req.Phone, req.Password, req.Confirmpass, req.IsEmployee, req.Role)
	return &v1.SignUpUserRespone{
		Id:         id,
		Phone:      phone,
		Password:   password,
		IsEmployee: isemployee,
		Role:       role,
	}, nil
}

func (s *UserService) SignInUser(ctx context.Context, req *v1.SignInUserRequest) (*v1.SignInUserRespone, error) {
	accessToken, refreshToken, err := s.uc.SignInUser(ctx, req.Phone, req.Password, req.Role)
	if err != nil {
		return nil, errors.New("invalid phone or password")
	}
	return &v1.SignInUserRespone{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) GetAllUser(ctx context.Context, req *v1.GetAllUserRequest) (*v1.GetAllUserRespone, error) {
	users, err := s.uc.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}
	var userList []*v1.User
	for _, user := range users {
		userList = append(userList, &v1.User{
			Id:         user.IDUser,
			Phone:      user.Phone,
			Password:   user.Password,
			Role:       user.Role,
			IsEmployee: user.IsEmployee,
		})
	}
	return &v1.GetAllUserRespone{Users: userList}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserRespone, error) {
	err := s.uc.DeleteUser(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserRespone{}, nil
}

func (s *UserService) RefreshAccessToken(ctx context.Context, req *v1.RefreshAccessTokenRequest) (*v1.RefreshAccessTokenRespone, error) {
	token, err := s.uc.RefreshAccessToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshAccessTokenRespone{AccessToken: token}, nil
}

func (s *UserService) GetUserByPage(ctx context.Context, req *v1.GetUserByPageRequest) (*v1.GetUsertByPageRespone, error) {
	user, err := s.uc.GetUserByPage(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	var users []*v1.User
	for _, us := range user {
		users = append(users, &v1.User{
			Id:         us.IDUser,
			Phone:      us.Phone,
			Password:   us.Password,
			Role:       us.Role,
			IsEmployee: us.IsEmployee,
		})
	}

	return &v1.GetUsertByPageRespone{Users: users}, err
}
