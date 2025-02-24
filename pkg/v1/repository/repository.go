package repository

import "frec.kr/tdoo/pkg/v1/gen/tdoo/orm"

type Repository struct {
	DB *orm.Client
}

func NewRepository(db *orm.Client) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) Close() error {
	return repo.DB.Close()
}

func (repo *Repository) Task() *TaskRepository {
	return NewTaskRepository(repo.DB)
}

func (repo *Repository) User() *UserRepository {
	return NewUserRepository(repo.DB)
}
