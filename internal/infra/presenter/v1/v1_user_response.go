package v1

import "simple-clean-architecture/internal/domain"

type UserCreateResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUserCreateResponse(user *domain.User) *UserCreateResponse {
	return &UserCreateResponse{
		ID:          user.ID,
		Title:       user.Title,
		Description: user.Description,
	}
}

type UserFindByIDResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUserFindByIDResponse(user *domain.User) *UserFindByIDResponse {
	return &UserFindByIDResponse{
		ID:          user.ID,
		Title:       user.Title,
		Description: user.Description,
	}
}

type UserListResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUserListResponse(users []*domain.User) []*UserListResponse {
	var userListResponse []*UserListResponse

	for _, user := range users {
		userListResponse = append(userListResponse, &UserListResponse{
			ID:          user.ID,
			Title:       user.Title,
			Description: user.Description,
		})
	}

	return userListResponse
}

type UserUpdateResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUserUpdateResponse(user *domain.User) *UserUpdateResponse {
	return &UserUpdateResponse{
		ID:          user.ID,
		Title:       user.Title,
		Description: user.Description,
	}
}
