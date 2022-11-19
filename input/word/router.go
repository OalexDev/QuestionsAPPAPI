package word

import (
	"github.com/OalexDev/QuestionsAPPAPI/exception"
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

	router.GET("/word/new", r.controller.GetNewWord)
	router.GET("/word", r.controller.GetWord)

}
