package room

import (
	"net/http"

	service "github.com/Five-Series/questions/app/room"
	"github.com/Five-Series/questions/exception"
	"github.com/Five-Series/questions/httphandler"
	"github.com/Five-Series/questions/input"
	"github.com/gin-gonic/gin"
)

const entity = "Room"

type Controller struct {
	input.DefaultController
	Service *service.Service
}

func NewControllerWord(svc *service.Service) *Controller {

	return &Controller{
		Service: svc,
		DefaultController: input.DefaultController{
			Environment: svc.Env,
		},
	}
}

func (c *Controller) GetRooms(context *gin.Context) {

	result, err := c.Service.GetRooms()
	if err != nil {
		switch err.(type) {

		case exception.InvalidParametersError:
			_ = httphandler.WriteInvalidParametersError(context, []string{err.Error()})
			return
		case exception.ConflictError:
			_ = httphandler.WriteConflictError(context, entity)
			return
		default:
			_ = httphandler.WriteInternalServerError(context, err.Error())
			return
		}
	}
	httphandler.WriteSuccess(context, http.StatusOK, entity, result)

}
