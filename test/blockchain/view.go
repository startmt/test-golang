package blockchain

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GeBlockChainView(c *fiber.Ctx, repository Collection) error {
	res, err := GetBlockChainArrayController(DefaultControllerResource{Context: c.Context(), Collection: repository})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetBlockChainByHashView(c *fiber.Ctx, repository Collection) error {
	hash := c.Params("hash")
	res, err := GetBlockChainByHashController(DefaultControllerResource{Context: c.Context(), Collection: repository}, hash)
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(err.Error())
		}

		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func GetBlockChainByIndex(c *fiber.Ctx, repository Collection) error {
	index, err := strconv.Atoi(c.Params("index"))
	if err != nil {
		return c.SendStatus(404)
	}

	res, err := GetBlockChainByIndexController(DefaultControllerResource{Context: c.Context(), Collection: repository}, index)
	return c.Status(fiber.StatusOK).JSON(res)
}

func AddBlockChainView(c *fiber.Ctx, repository Collection) error {
	if err := AddBlockChainController(DefaultControllerResource{Context: c.Context(), Collection: repository}, c.Body()); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.SendStatus(fiber.StatusCreated)
}

func ValidateBlockView(c *fiber.Ctx, repository Collection) error {
	res, err := ValidateBlockChainController(DefaultControllerResource{Context: c.Context(), Collection: repository})
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
