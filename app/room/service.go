package room

import (
	"github.com/Five-Series/questions/app/room/model"
	"github.com/Five-Series/questions/infra/environment"
	outRoom "github.com/Five-Series/questions/output/room"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Service struct {
	AWSSess    *session.Session
	Env        *environment.Environment
	Repository *outRoom.Repository
}

func NewService(env *environment.Environment, sess *session.Session, repo *outRoom.Repository) *Service {
	return &Service{
		Env:        env,
		AWSSess:    sess,
		Repository: repo,
	}
}

func (s Service) GetRooms() ([]model.Rooms, error) {

	rooms, err := s.Repository.GetRooms()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
