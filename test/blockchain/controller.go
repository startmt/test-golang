package blockchain

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBlockChainArrayController(c *fiber.Ctx, repository Collection) error {
	blocks, err := GetAllBlockChain(repository)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.JSON(blocks)
}

func GetBlockChainByHashController(c *fiber.Ctx, repository Collection) error {
	searchChain, err := GetBlockChainBy(QueryOneBlockChainByHash(c.Params("hash")))(repository)
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(searchChain)
}

func GetBlockChainByIndexController(c *fiber.Ctx, repository Collection) error {
	index, err := strconv.Atoi(c.Params("index"))
	if err != nil {
		return c.SendStatus(404)
	}

	searchChain, err := GetBlockChainBy(QueryOneBlockChainByIndex(index))(repository)
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			return c.SendStatus(404)
		}
		return c.SendStatus(400)
	}
	return c.JSON(searchChain)
}

func AddBlockChainController(c *fiber.Ctx, repository Collection) error {
	var reqBody CreateBlockChainRequest
	if err := json.Unmarshal(c.Body(), &reqBody); err != nil {
		return c.SendStatus(400)
	}

	blocks, err := GetAllBlockChain(repository)
	if err != nil {
		return c.SendStatus(400)
	}

	block := BlockChain{
		Index: len(blocks),
		Body:  reqBody.Body,
	}

	serviceParam := AddPrevHashInBlock(blocks, block)
	newBlock := NewBlockByBody(serviceParam)

	if err = AppendBlockInDatabase(repository, newBlock); err != nil {
		return c.SendStatus(400)
	}

	return c.SendStatus(201)
}

func ValidateBlockChainController(c *fiber.Ctx, repository Collection) error {
	blocks, err := GetAllBlockChain(repository)
	if err != nil {
		return c.SendStatus(400)
	}

	isBlockValidate := ValidateBlockChain(blocks)
	response := ValidateBlockChainResponse{IsValidate: isBlockValidate}

	return c.Status(fiber.StatusOK).JSON(response)
}
