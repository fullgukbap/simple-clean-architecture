package utils

import (
	"simple-clean-architecture/internal/domain"
	"simple-clean-architecture/internal/repository/ent"
)

func ConvertDomainUser(physicalUser *ent.User) *domain.User {
	return &domain.User{
		ID:          physicalUser.ID,
		Title:       physicalUser.Title,
		Description: physicalUser.Description,
	}
}

func ConvertDomainUsers(physicalUsers []*ent.User) []*domain.User {
	var users []*domain.User
	for _, pysicalUser := range physicalUsers {
		users = append(users, ConvertDomainUser(pysicalUser))
	}

	return users
}
