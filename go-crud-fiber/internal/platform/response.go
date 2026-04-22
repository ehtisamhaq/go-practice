package platform

import "github.com/gofiber/fiber/v3"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(c fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c fiber.Ctx, status int, message string, err interface{}) error {
	return c.Status(status).JSON(APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func ErrorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ErrorResponse(c, code, "An error occurred", err.Error())
}
