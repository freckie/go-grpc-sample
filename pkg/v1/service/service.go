package service

import "frec.kr/tdoo/pkg/v1/repository"

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Task() *TaskService {
	return NewTaskService(
		s.repo.Task(),
		s.repo.User(),
	)
}

func (s *Service) User() *UserService {
	return NewUserService(
		s.repo.User(),
	)
}
