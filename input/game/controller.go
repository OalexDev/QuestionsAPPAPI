package game

import (
	"net/http"
	"strconv"

	service "github.com/OalexDev/QuestionsAPPAPI/app/game"
	"github.com/OalexDev/QuestionsAPPAPI/app/game/model"
	"github.com/OalexDev/QuestionsAPPAPI/exception"
	"github.com/OalexDev/QuestionsAPPAPI/httphandler"
	"github.com/OalexDev/QuestionsAPPAPI/input"
	"github.com/gin-gonic/gin"
)

const entity = "Game"

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

	room := context.GetHeader("X-ROOM-ID")
	if len(room) == 0 {
		err := exception.NewInvalidParametersError([]string{"X-ROOM-ID"})
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}
	userID := context.GetHeader("X-PLAYER-ID")
	if len(userID) == 0 {
		err := exception.NewInvalidParametersError([]string{"X-PLAYER-ID"})
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}

	roomID, err := strconv.ParseInt(room, 10, 64)
	if err != nil {
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}
	playerID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}

	gamePlay := model.Game{}
	err = context.BindJSON(&gamePlay)
	if err != nil {
		httphandler.WriteInvalidParametersError(context, []string{"text"})
		return
	}

	gamePlay.RoomID = roomID
	gamePlay.PlayerID = playerID

	err = c.Service.InsertMaessage(&gamePlay)
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

	httphandler.WriteSuccess(context, http.StatusCreated, entity, gamePlay.Text)
}

func (c *Controller) GetGameMessages(context *gin.Context) {

	room := context.GetHeader("X-ROOM-ID")
	if len(room) == 0 {
		err := exception.NewInvalidParametersError([]string{"X-ROOM-ID"})
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}
	roomID, err := strconv.ParseInt(room, 10, 64)
	if err != nil {
		_ = httphandler.WriteMissingParametersError(context, []string{err.Error()})
		return
	}

	result, err := c.Service.GetGameMessages(roomID)
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
