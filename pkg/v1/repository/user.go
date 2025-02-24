package repository

import (
	"context"

	"frec.kr/tdoo/pkg/v1/gen/tdoo"
	"frec.kr/tdoo/pkg/v1/gen/tdoo/orm"
)

type UserRepository struct {
	DB *orm.Client
}

func NewUserRepository(db *orm.Client) *UserRepository {
	return &UserRepository{DB: db}
}

func (r UserRepository) GetById(ctx context.Context, id []byte) (*tdoo.User, error) {
	// TODO implement me
	panic("implement me")
}
