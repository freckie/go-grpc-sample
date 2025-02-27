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

func (t TaskService) Get(ctx context.Context, req *tdoo.GetTaskRequest) (*tdoo.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) List(ctx context.Context, req *tdoo.ListTaskRequest) (*tdoo.ListTaskResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) Create(ctx context.Context, req *tdoo.CreateTaskRequest) (*tdoo.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) Update(ctx context.Context, req *tdoo.UpdateTaskRequest) (*tdoo.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) Delete(ctx context.Context, req *tdoo.DeleteTaskRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

func (t TaskService) Shift(ctx context.Context, req *tdoo.ShiftTaskRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}
