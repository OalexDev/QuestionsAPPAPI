package healthcheck

import (
	"net/http"

	service "github.com/OalexDev/QuestionsAPPAPI/app/healthcheck"
	"github.com/OalexDev/QuestionsAPPAPI/exception"
	"github.com/OalexDev/QuestionsAPPAPI/httphandler"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *service.Service
}

func NewController(svc *service.Service) *Controller {

	return &Controller{
		Service: svc,
	}

}

const entity = "health"

// Health Responsible for testing life program
func (c *Controller) GetHealth(context *gin.Context) {

	result, err := c.Service.Get()
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

	err = httphandler.WriteSuccess(context, http.StatusOK, entity, result)

}
