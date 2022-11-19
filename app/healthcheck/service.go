package healthcheck

import (
	"github.com/OalexDev/QuestionsAPPAPI/app/healthcheck/model"
	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"
	outHealthCheck "github.com/OalexDev/QuestionsAPPAPI/output/healthcheck"
)

type Service struct {
	Env        *environment.Environment
	Repository *outHealthCheck.Repository
}

func NewService(repo *outHealthCheck.Repository, env *environment.Environment) *Service {
	return &Service{
		Env:        env,
		Repository: repo,
	}
}

func (s Service) Get() (model.Health, error) {
	h := model.Health{}
	err := s.Repository.Read(&h)
	return h, err
}
