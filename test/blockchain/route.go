package blockchain

import (
	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router, repository Collection) {
	router.Get("/blockchain", InjectBlockchainRepositoryIntoController(repository, GeBlockChainView))
	router.Get("/blockchain/hash/:hash", InjectBlockchainRepositoryIntoController(repository, GetBlockChainByHashView))
	router.Get("/blockchain/index/:index", InjectBlockchainRepositoryIntoController(repository, GetBlockChainByIndex))
	router.Get("/blockchain/validate", InjectBlockchainRepositoryIntoController(repository, ValidateBlockView))
	router.Post("/blockchain", InjectBlockchainRepositoryIntoController(repository, AddBlockChainView))
}
