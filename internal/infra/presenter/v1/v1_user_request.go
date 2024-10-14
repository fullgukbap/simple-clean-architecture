package v1

import "simple-clean-architecture/internal/domain"

type UserCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r UserCreateRequest) ToDodmain() *domain.User {
	return &domain.User{
		Title:       r.Title,
		Description: r.Description,
	}
}

type UserUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r UserUpdateRequest) ToDomain() *domain.User {
	return &domain.User{
		Title:       r.Title,
		Description: r.Description,
	}
}
