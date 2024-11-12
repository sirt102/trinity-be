//go:build wireinject

package wire

import (
	"trinity-be/internal/handlers"
	"trinity-be/internal/repositories"
	"trinity-be/internal/services"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*handlers.UserHandler) {
    wire.Build(
        repositories.NewUserRepository,
        services.NewUserService,
        handlers.NewUserHandler,
    )
    return new(handlers.UserHandler)
}