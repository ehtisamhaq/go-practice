// internal/user/handler.go
package user

import (
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

func RegisterRoutes(group fiber.Router, service *UserService) {
	handler := &UserHandler{service: service, v: validator.New()}

	group.Post("/users", handler.Register)
	// Add more routes for other CRUD operations (e.g., GET, PUT, DELETE)
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	req := new(RegisterUserDTO)

	// Fiber v3 Unified Binding: Checks Body, Query, and Params automatically
	if err := c.Bind().All(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validate the request
	if err := h.v.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully", "data": req})
}
