package word

import (
	"fmt"
	"net/http"
	"strconv"

	service "github.com/OalexDev/QuestionsAPPAPI/app/word"
	"github.com/OalexDev/QuestionsAPPAPI/exception"
	"github.com/OalexDev/QuestionsAPPAPI/httphandler"
	"github.com/OalexDev/QuestionsAPPAPI/input"
	"github.com/gin-gonic/gin"
)

const entity = "Word"

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

func (c *Controller) GetNewWord(context *gin.Context) {

	roomID := context.GetHeader("X-ROOM-ID")
	if len(roomID) == 0 {
		err := exception.NewInvalidParametersError([]string{"X-ROOM-ID"})
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}

	roomInt, err := strconv.ParseInt(roomID, 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing room ID: %v", err)
		httphandler.WriteSuccess(context, http.StatusConflict, entity, err)
		return
	}

	result, err := c.Service.GetNewWord(roomInt)
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

func (c *Controller) GetWord(context *gin.Context) {

	roomID := context.GetHeader("X-ROOM-ID")
	if len(roomID) == 0 {
		err := exception.NewInvalidParametersError([]string{"X-ROOM-ID"})
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}

	roomInt, err := strconv.ParseInt(roomID, 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing room ID: %v", err)
		httphandler.WriteSuccess(context, http.StatusConflict, entity, err)
		return
	}

	result, err := c.Service.GetWord(roomInt)
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
