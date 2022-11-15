package game

import (
	"github.com/Five-Series/questions/exception"
	"github.com/gin-gonic/gin"
)

type Router struct {
	controller *Controller
}

func NewRouter(controller *Controller) (*Router, error) {

	if controller == nil {
		return nil, exception.NewInvalidParametersError([]string{"controller"})
	}

	return &Router{
		controller: controller,
	}, nil
}

// SetRoutes resposible for create router on endpoint group
func (r *Router) SetRoutes(router *gin.RouterGroup) {

	router.POST("/game", r.controller.GetRooms)
	router.GET("/game/messages", r.controller.GetGameMessages)

	// Alexandre Permitir Votar qual o melhor, mas nao pode votar em vc mesmo
}
