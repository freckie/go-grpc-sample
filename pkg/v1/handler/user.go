package handler

import (
	"context"

	"frec.kr/tdoo/pkg/v1/gen/tdoo"
	"frec.kr/tdoo/pkg/v1/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	svc *service.UserService

	tdoo.UnimplementedUserServiceServer
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (u UserHandler) Get(ctx context.Context, req *tdoo.GetUserRequest) (*tdoo.User, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (u UserHandler) List(ctx context.Context, req *tdoo.ListUserRequest) (*tdoo.ListUserResponse, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (u UserHandler) Create(ctx context.Context, req *tdoo.CreateUserRequest) (*tdoo.User, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (u UserHandler) Update(ctx context.Context, req *tdoo.UpdateUserRequest) (*tdoo.User, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (u UserHandler) Delete(ctx context.Context, req *tdoo.DeleteUserRequest) (*emptypb.Empty, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}
