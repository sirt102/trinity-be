//go:build wireinject

package wire

import (
	"trinity-be/internal/handlers"
	"trinity-be/internal/repositories"
	"trinity-be/internal/services"

	"github.com/google/wire"
)

func InitVoucherRouterHandler() *handlers.VoucherHandler {
	wire.Build(
		repositories.NewVoucherRepository,
		services.NewVoucherService,
		handlers.NewVoucherHandler,
	)

	return new(handlers.VoucherHandler)
}
