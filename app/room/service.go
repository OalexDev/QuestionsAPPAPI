package room

import (
	"fmt"

	"github.com/OalexDev/QuestionsAPPAPI/app/room/model"
	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"
	outRoom "github.com/OalexDev/QuestionsAPPAPI/output/room"
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

func (s Service) EntryRoom(room *model.Rooms) (string, error) {

	id, err := s.Repository.EntryRoom(room)
	if err != nil {
		return "", err
	}
	if id < 1 {
		return "", fmt.Errorf("no insert success: %d", err)
	}

	return "Success", nil
}

func (s Service) MakeRoom() (int64, error) {

	id, err := s.Repository.MakeRoom()
	if err != nil {
		return 0, err
	}
	if id < 1 {
		return 0, fmt.Errorf("no insert success: %d", err)
	}

	return id, nil
}
