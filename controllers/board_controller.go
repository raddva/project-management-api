package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raddva/projeqtor-api-go/models"
	"github.com/raddva/projeqtor-api-go/services"
	"github.com/raddva/projeqtor-api-go/utils"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{service: s}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	board := new(models.Board)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx, "Failed to read Request", err.Error())
	}

	if err := c.service.Create(board); err != nil {
		return utils.BadRequest(ctx, "Failed Saving Data", err.Error())
	}

	return utils.Success(ctx, "Board Created Successfully", board)
}