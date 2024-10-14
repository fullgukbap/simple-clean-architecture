package repository

import (
	"context"
	"simple-clean-architecture/internal/domain"
	"simple-clean-architecture/internal/repository/database"
	"simple-clean-architecture/internal/repository/utils"
)

var _ domain.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	*database.Database
}

func NewUserRepository(database *database.Database) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser, err := r.Client.User.
		Create().
		SetTitle(user.Title).
		SetDescription(user.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUser(createdUser), nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	findedUser, err := r.Client.User.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUser(findedUser), nil
}

func (r *UserRepository) List(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	listUsers, err := r.Client.User.
		Query().
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUsers(listUsers), nil
}

func (r *UserRepository) Update(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
	updatedUser, err := r.Client.User.
		UpdateOneID(id).
		SetTitle(user.Title).
		SetDescription(user.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUser(updatedUser), nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	err := r.Client.User.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
