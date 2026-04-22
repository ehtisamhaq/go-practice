// internal/user/handler.go
package user

import (
	"fmt"
	"go-crud-fiber/internal/platform"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service *UserService
	v       *validator.Validate
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{service: s, v: validator.New()}
}

func RegisterRoutes(group fiber.Router, handler *UserHandler) {
	group.Post("/", handler.Register)
	group.Get("/", handler.GetAll)
	group.Get("/:id", handler.GetByID)
	group.Put("/:id", handler.Update)
	group.Delete("/:id", handler.Delete)
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	req := new(RegisterUserDTO)

	if err := c.Bind().All(req); err != nil {
		return platform.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := h.v.Struct(req); err != nil {
		return platform.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed", err.Error())
	}

	userModel := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.service.Register(userModel); err != nil {
		return platform.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to register user", err.Error())
	}

	return platform.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", userModel)
}

func (h *UserHandler) GetAll(c fiber.Ctx) error {
	users, err := h.service.GetAll()
	if err != nil {
		return platform.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch users", err.Error())
	}
	return platform.SuccessResponse(c, fiber.StatusOK, "Users fetched successfully", users)
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	idStr := c.Params("id")
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return platform.ErrorResponse(c, fiber.StatusBadRequest, "Invalid ID format", nil)
	}
	user, err := h.service.GetByID(id)
	if err != nil {
		return platform.ErrorResponse(c, fiber.StatusNotFound, "User not found", nil)
	}
	return platform.SuccessResponse(c, fiber.StatusOK, "User found", user)
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	idStr := c.Params("id")
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return platform.ErrorResponse(c, fiber.StatusBadRequest, "Invalid ID format", nil)
	}
	req := new(RegisterUserDTO)
	if err := c.Bind().All(req); err != nil {
		return platform.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		return platform.ErrorResponse(c, fiber.StatusNotFound, "User not found", nil)
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password

	if err := h.service.Update(user); err != nil {
		return platform.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return platform.SuccessResponse(c, fiber.StatusOK, "User updated successfully", user)
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
	idStr := c.Params("id")
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return platform.ErrorResponse(c, fiber.StatusBadRequest, "Invalid ID format", nil)
	}
	if err := h.service.Delete(id); err != nil {
		return platform.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete user", err.Error())
	}
	return platform.SuccessResponse(c, fiber.StatusOK, "User deleted successfully", nil)
}
