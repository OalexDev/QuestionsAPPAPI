package healthcheck

import (
	"database/sql"

	"github.com/Five-Series/questions/app/healthcheck/model"
	"github.com/gin-gonic/gin"
)

type Repository struct {
	DB         *sql.DB
	RoutesInfo []gin.RouteInfo
}

func NewRepository(db *sql.DB, routes []gin.RouteInfo) *Repository {
	return &Repository{
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
		Method:  "GET",
		Path:    "/health",
		Message: "you are here",
	})

	return nil
}
