package healthcheck

import (
	"database/sql"

	"github.com/Five-Series/questions/app/healthcheck/model"
	"github.com/Five-Series/questions/infra/environment"
	"github.com/gin-gonic/gin"
)

type Repository struct {
	Env        *environment.Environment
	DB         *sql.DB
	RoutesInfo []gin.RouteInfo
}

func NewRepository(db *sql.DB, routes []gin.RouteInfo, env *environment.Environment) *Repository {
	return &Repository{
		Env:        env,
		DB:         db,
		RoutesInfo: routes,
	}
}

func (r Repository) Read(health *model.Health) error {

	health.Message = "alive and kicking"

	for _, ri := range r.RoutesInfo {
		health.Routes = append(health.Routes, model.Route{
			Method: ri.Method,
			Path:   ri.Path,
		})
	}

	health.Routes = append(health.Routes, model.Route{
		Method:      "GET",
		Path:        "/health",
		Message:     "you are here",
		Version:     r.Env.App.App_version,
		Environment: r.Env.Env,
	})

	return nil
}
