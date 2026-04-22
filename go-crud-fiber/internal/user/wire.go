package user

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserModule(db *gorm.DB) *UserHandler {
	wire.Build(
		NewRepository,
		NewUserService,
		NewUserHandler,
	)
	return &UserHandler{}
}
