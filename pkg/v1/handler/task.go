package handler

import (
	"context"

	"frec.kr/tdoo/pkg/v1/gen/tdoo"
	"frec.kr/tdoo/pkg/v1/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TaskHandler struct {
	svc *service.TaskService

	tdoo.UnimplementedTaskServiceServer
}

func NewTaskHandler(svc *service.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

func (t TaskHandler) Get(ctx context.Context, req *tdoo.GetTaskRequest) (*tdoo.Task, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (t TaskHandler) List(ctx context.Context, req *tdoo.ListTaskRequest) (*tdoo.ListTaskResponse, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (t TaskHandler) Create(ctx context.Context, req *tdoo.CreateTaskRequest) (*tdoo.Task, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (t TaskHandler) Update(ctx context.Context, req *tdoo.UpdateTaskRequest) (*tdoo.Task, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (t TaskHandler) Delete(ctx context.Context, req *tdoo.DeleteTaskRequest) (*emptypb.Empty, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (t TaskHandler) Shift(ctx context.Context, req *tdoo.ShiftTaskRequest) (*emptypb.Empty, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}
