package healthcheck

import (
	"database/sql"
	"log"

	health "github.com/Five-Series/questions/app/healthcheck"
	inHealthCheck "github.com/Five-Series/questions/input/healthcheck"
	outHealthCheck "github.com/Five-Series/questions/output/healthcheck"
	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	DbConnection *sql.DB
	RouterGroup  *gin.RouterGroup
	// Env          *environment.Environment
	RoutesGin []gin.RouteInfo
}

func (e *Healthcheck) Start() error {

	repo := outHealthCheck.NewRepository(e.DbConnection, e.RoutesGin)

	svc := health.NewService(repo)

	ctl := inHealthCheck.NewController(svc)

	router, err := inHealthCheck.NewRouter(ctl)
	if err != nil {
		log.Fatalf("error on build router.: %v", err)
	}
	router.SetRoutes(e.RouterGroup)

	return nil

}
