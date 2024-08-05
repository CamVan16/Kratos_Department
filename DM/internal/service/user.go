package service

import (
	v1 "DM/api/helloworld/v1"
	"DM/internal/biz"
	"context"
	"errors"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	biz *biz.UserService
}

func NewUserService(biz *biz.UserService) *UserService {
	return &UserService{biz: biz}
}

func (s *UserService) SignUpUser(ctx context.Context, req *v1.SignUpUserRequest) (*v1.SignUpUserRespone, error) {
	id, phone, password, isemployee, role := s.biz.SignUpUser(ctx, req.Phone, req.Password, req.Confirmpass, req.IsEmployee, req.Role)
	return &v1.SignUpUserRespone{
		Id:         id,
		Phone:      phone,
		Password:   password,
		IsEmployee: isemployee,
		Role:       role,
	}, nil
}

func (s *UserService) SignInUser(ctx context.Context, req *v1.SignInUserRequest) (*v1.SignInUserRespone, error) {
	accessToken, refreshToken, err := s.biz.SignInUser(ctx, req.Phone, req.Password, req.Role)
	if err != nil {
		return nil, errors.New("invalid phone or password")
	}
	return &v1.SignInUserRespone{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) GetAllUser(ctx context.Context, req *v1.GetAllUserRequest) (*v1.GetAllUserRespone, error) {
	users, err := s.biz.GetAllUser(ctx)
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
	err := s.biz.DeleteUser(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserRespone{}, nil
}

func (s *UserService) RefreshAccessToken(ctx context.Context, req *v1.RefreshAccessTokenRequest) (*v1.RefreshAccessTokenRespone, error) {
	token, err := s.biz.RefreshAccessToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshAccessTokenRespone{AccessToken: token}, nil
}
