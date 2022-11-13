package healthcheck

import (
	"github.com/Five-Series/questions/app/healthcheck/model"
	outHealthCheck "github.com/Five-Series/questions/output/healthcheck"
)

type Service struct {
	// HealthReader model.Health
	Repository *outHealthCheck.Repository
}

func NewService(repo *outHealthCheck.Repository) *Service {
	return &Service{
		// HealthReader: reader,
		Repository: repo,
	}
}

func (s Service) Get() (model.Health, error) {
	h := model.Health{}
	err := s.Repository.Read(&h)
	return h, err
}
