package blockchain

import (
	"github.com/gofiber/fiber/v2"
)

func InjectBlockchainRepositoryIntoController(repository Collection, controllerFn func(*fiber.Ctx, Collection) error) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return controllerFn(ctx, repository)
	}
}
