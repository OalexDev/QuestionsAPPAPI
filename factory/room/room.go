package room

import (
	"database/sql"
	"log"

	svcRoom "github.com/Five-Series/questions/app/room"
	"github.com/Five-Series/questions/infra/environment"
	inRoom "github.com/Five-Series/questions/input/room"
	outRoom "github.com/Five-Series/questions/output/room"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

type Room struct {
	DbConnection *sql.DB

	RouterGroup *gin.RouterGroup
	Env         *environment.Environment
	AWSSess     *session.Session
}

func (e *Room) Start() error {

	repo := outRoom.NewRepository(
		e.Env,
		e.DbConnection,
	)

	svc := svcRoom.NewService(e.Env, e.AWSSess, repo)

	ctl := inRoom.NewControllerWord(svc)

	room, err := inRoom.NewRouter(ctl)
	if err != nil {
		log.Fatalf("error on build router.: %v", err)
	}
	room.SetRoutes(e.RouterGroup)

	return nil

}
