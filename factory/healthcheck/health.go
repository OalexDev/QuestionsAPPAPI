package healthcheck

import (
	"database/sql"
	"log"

	health "github.com/OalexDev/QuestionsAPPAPI/app/healthcheck"
	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"
	inHealthCheck "github.com/OalexDev/QuestionsAPPAPI/input/healthcheck"
	outHealthCheck "github.com/OalexDev/QuestionsAPPAPI/output/healthcheck"
	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	DbConnection *sql.DB
	RouterGroup  *gin.RouterGroup
	Env          *environment.Environment
	RoutesGin    []gin.RouteInfo
}

func (e *Healthcheck) Start() error {

	repo := outHealthCheck.NewRepository(e.DbConnection, e.RoutesGin, e.Env)

	svc := health.NewService(repo, e.Env)

	ctl := inHealthCheck.NewController(svc)

	router, err := inHealthCheck.NewRouter(ctl)
	if err != nil {
		log.Fatalf("error on build router.: %v", err)
	}
	router.SetRoutes(e.RouterGroup)

	return nil

}
