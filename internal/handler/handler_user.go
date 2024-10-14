package handler

import (
	"simple-clean-architecture/internal/domain"
	v1 "simple-clean-architecture/internal/infra/presenter/v1"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecae domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecae,
	}
}

func (r *UserHandler) Route(router fiber.Router) {
	router.Post("/users", r.Create)
	router.Get("/users", r.List)
	router.Get("/users/:id", r.FindByID)
	router.Put("/users/:id", r.Update)
	router.Delete("/users/:id", r.Delete)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	body := new(v1.UserCreateRequest)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	createdUser, err := h.userUsecase.Create(c.Context(), body.ToDodmain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(v1.NewUserCreateResponse(createdUser))
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	offset := c.QueryInt("offset")
	limit := c.QueryInt("limit")

	listUsers, err := h.userUsecase.List(c.Context(), offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewUserListResponse(listUsers))
}

func (h *UserHandler) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	findedUser, err := h.userUsecase.FindByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewUserFindByIDResponse(findedUser))
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	body := new(v1.UserUpdateRequest)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	updatedUser, err := h.userUsecase.Update(c.Context(), id, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewUserUpdateResponse(updatedUser))
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = h.userUsecase.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
