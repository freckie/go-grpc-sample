package repository

import (
	"context"

	"frec.kr/tdoo/pkg/v1/gen/tdoo/orm"
)

type TaskRepository struct {
	DB *orm.Client
}

func NewTaskRepository(db *orm.Client) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r TaskRepository) GetById(ctx context.Context, id []byte) (*orm.Task, error) {
	// TODO implement me
	panic("implement me")
}
