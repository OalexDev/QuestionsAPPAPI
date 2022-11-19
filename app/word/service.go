package word

import (
	model "github.com/OalexDev/QuestionsAPPAPI/app/word/model"
	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"
	outword "github.com/OalexDev/QuestionsAPPAPI/output/word"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Service struct {
	AWSSess    *session.Session
	Env        *environment.Environment
	Repository *outword.Repository
}

func NewService(env *environment.Environment, sess *session.Session, repo *outword.Repository) *Service {
	return &Service{
		Env:        env,
		AWSSess:    sess,
		Repository: repo,
	}
}

func (s Service) GetNewWord(room int64) (model.Word, error) {

	word, err := s.Repository.ReadeOneNewWord(room)
	if err != nil {
		return word, err
	}

	_, err = s.Repository.InsertWordToRoom(room, word.ID)
	if err != nil {
		return word, err
	}

	return word, nil
}

func (s Service) GetWord(room int64) (model.Word, error) {

	word, err := s.Repository.ReadWord(room)
	if err != nil {
		return word, err
	}

	return word, nil
}
