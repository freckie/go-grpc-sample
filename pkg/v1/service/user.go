package service

import (
	"context"

	"frec.kr/tdoo/pkg/v1/gen/tdoo"
	"frec.kr/tdoo/pkg/v1/repository"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	user *repository.UserRepository

	tdoo.UserServiceServer
}

func NewUserService(user *repository.UserRepository) *UserService {
	return &UserService{
		user: user,
	}
}

func (u UserService) Get(ctx context.Context, req *tdoo.GetUserRequest) (*tdoo.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserService) List(ctx context.Context, req *tdoo.ListUserRequest) (*tdoo.ListUserResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserService) Create(ctx context.Context, req *tdoo.CreateUserRequest) (*tdoo.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserService) Update(ctx context.Context, req *tdoo.UpdateUserRequest) (*tdoo.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserService) Delete(ctx context.Context, req *tdoo.DeleteUserRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}
