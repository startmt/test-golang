package blockchain

import (
	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router, repository Collection) {
	router.Get("/blockchain", InjectBlockchainRepositoryIntoController(repository, GetBlockChainArrayController))
	router.Get("/blockchain/hash/:hash", InjectBlockchainRepositoryIntoController(repository, GetBlockChainByHashController))
	router.Get("/blockchain/index/:index", InjectBlockchainRepositoryIntoController(repository, GetBlockChainByIndexController))
	router.Get("/blockchain/validate", InjectBlockchainRepositoryIntoController(repository, ValidateBlockChainController))
	router.Post("/blockchain", InjectBlockchainRepositoryIntoController(repository, AddBlockChainController))
}
