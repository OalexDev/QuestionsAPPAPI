package word

import (
	model "github.com/Five-Series/questions/app/word/model"
	"github.com/Five-Series/questions/infra/environment"
	outword "github.com/Five-Series/questions/output/word"
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

func (s Service) GetWord() ([]model.Word, error) {

	word, err := s.Repository.ReadeOneWord()
	if err != nil {
		return nil, err
	}

	return word, nil
}
