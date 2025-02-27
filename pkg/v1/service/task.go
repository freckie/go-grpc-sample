package service

import (
	"context"

	"frec.kr/tdoo/pkg/v1/gen/tdoo"
	"frec.kr/tdoo/pkg/v1/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TaskService struct {
	task *repository.TaskRepository
	user *repository.UserRepository

	tdoo.TaskServiceServer
}

func NewTaskService(
	task *repository.TaskRepository,
	user *repository.UserRepository,
) *TaskService {
	return &TaskService{
		task: task,
		user: user,
	}
}

func (t TaskService) Get(ctx context.Context, req *tdoo.TaskGetRequest) (*tdoo.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) List(ctx context.Context, req *tdoo.TaskListRequest) (*tdoo.TaskListResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) Create(ctx context.Context, req *tdoo.TaskCreateRequest) (*tdoo.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) Update(ctx context.Context, req *tdoo.TaskUpdateRequest) (*tdoo.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) Delete(ctx context.Context, req *tdoo.TaskDeleteRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) Shift(ctx context.Context, req *tdoo.TaskShiftRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}
