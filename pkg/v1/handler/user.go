package handler

import (
	"context"

	"frec.kr/tdoo/pkg/v1/gen/tdoo"
	"frec.kr/tdoo/pkg/v1/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	svc *service.UserService

	tdoo.UnimplementedUserServiceServer
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (u UserHandler) Get(ctx context.Context, req *tdoo.UserGetRequest) (*tdoo.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) List(ctx context.Context, req *tdoo.UserListRequest) (*tdoo.UserListResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) Create(ctx context.Context, req *tdoo.UserCreateRequest) (*tdoo.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) Update(ctx context.Context, req *tdoo.UserUpdateRequest) (*tdoo.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) Delete(ctx context.Context, req *tdoo.UserDeleteRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}
