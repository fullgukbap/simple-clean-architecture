package usecase

import (
	"context"
	"simple-clean-architecture/internal/domain"
)

var _ domain.UserUsecase = (*UserUsecase)(nil)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.userRepository.Create(ctx, user)
}

func (u *UserUsecase) FindByID(ctx context.Context, id int) (*domain.User, error) {
	return u.userRepository.FindByID(ctx, id)
}
func (u *UserUsecase) List(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	return u.userRepository.List(ctx, offset, limit)
}

func (u *UserUsecase) Update(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
	return u.userRepository.Update(ctx, id, user)
}

func (u *UserUsecase) Delete(ctx context.Context, id int) error {
	return u.userRepository.Delete(ctx, id)
}
