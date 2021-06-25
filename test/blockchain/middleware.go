package blockchain

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func InjectBlockchainRepositoryIntoController(repository Collection, controllerFn func(*fiber.Ctx, Collection) error) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		err := controllerFn(ctx, repository)
		if err != nil {
			log.Println(err)
		}
		return err
	}
}
