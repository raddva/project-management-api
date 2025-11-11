package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/raddva/projeqtor-api-go/models"
	"github.com/raddva/projeqtor-api-go/repositories"
)

type BoardService interface {
	Create(board *models.Board) error
}

type boardService struct {
	boardRepo repositories.BoardRepository
	userRepo repositories.UserRepository	
}

func NeewBoardService(boardRepo repositories.BoardRepository, userRepo repositories.UserRepository) BoardService {
	return &boardService{boardRepo, userRepo}
}

func (s *boardService) Create (board *models.Board) error {
	user, err := s.userRepo.FindByPublicID(board.OwnerPublicID.String())
	if err != nil {
		return errors.New("user not found")
	}
	board.PublicID = uuid.New()
	board.OwnerID = user.InternalID
	return s.boardRepo.Create(board)

}