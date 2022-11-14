package room

import (
	"net/http"
	"strconv"

	service "github.com/Five-Series/questions/app/room"
	"github.com/Five-Series/questions/app/room/model"
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

	// roomID := context.GetHeader("X-ROOM-ID")
	// if len(roomID) == 0 {
	// 	err := exception.NewInvalidParametersError([]string{"X-ROOM-ID"})
	// 	_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
	// 	return
	// }
	// roomInt, err := strconv.ParseInt(roomID, 10, 64)
	// if err != nil {
	// 	err = fmt.Errorf("error parsing room ID: %v", err)
	// 	httphandler.WriteSuccess(context, http.StatusConflict, entity, err)
	// 	return
	// }

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

func (c *Controller) EntryRoom(context *gin.Context) {

	id := context.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if len(id) == 0 || err != nil {
		err := exception.NewInvalidParametersError([]string{"id"})
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}

	userID := context.GetHeader("X-USER-ID")
	if len(userID) == 0 {
		err := exception.NewInvalidParametersError([]string{"X-USER-ID"})
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}
	userInt, err := strconv.ParseInt(userID, 10, 64)
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

	room := model.Rooms{
		ID:     idInt,
		UserID: userInt,
	}

	result, err := c.Service.EntryRoom(&room)
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

	httphandler.WriteSuccess(context, http.StatusCreated, entity, result)

}

func (c *Controller) MakeRoom(context *gin.Context) {

	result, err := c.Service.MakeRoom()
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
	httphandler.WriteSuccess(context, http.StatusCreated, entity, result)

}
