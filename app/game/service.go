package game

import (
	"strings"

	"github.com/Five-Series/questions/app/game/model"
	"github.com/Five-Series/questions/infra/environment"
	outGame "github.com/Five-Series/questions/output/game"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Service struct {
	AWSSess    *session.Session
	Env        *environment.Environment
	Repository *outGame.Repository
}

func NewService(env *environment.Environment, sess *session.Session, repo *outGame.Repository) *Service {
	return &Service{
		Env:        env,
		AWSSess:    sess,
		Repository: repo,
	}
}

func (s Service) InsertMaessage(gamePlay *model.Game) error {

	word, err := s.Repository.GetWordByID(gamePlay.RoomID)
	if err != nil {
		return err
	}

	gamePlay.Text = strings.Replace(gamePlay.Text, "#", *word, -1)

	s.Repository.InsertMaessage(gamePlay)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetGameMessages(id int64) ([]model.Game, error) {

	mgs, err := s.Repository.GetGameMessages(id)
	if err != nil {
		return nil, err
	}

	return mgs, nil
}
