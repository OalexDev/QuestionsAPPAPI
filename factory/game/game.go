package game

import (
	"database/sql"
	"log"

	svcGame "github.com/OalexDev/QuestionsAPPAPI/app/game"

	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"
	inGame "github.com/OalexDev/QuestionsAPPAPI/input/game"
	outGame "github.com/OalexDev/QuestionsAPPAPI/output/game"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

type Game struct {
	DbConnection *sql.DB

	RouterGroup *gin.RouterGroup
	Env         *environment.Environment
	AWSSess     *session.Session
}

func (e *Game) Start() error {

	repo := outGame.NewRepository(
		e.Env,
		e.DbConnection,
	)

	svc := svcGame.NewService(e.Env, e.AWSSess, repo)

	ctl := inGame.NewControllerWord(svc)

	room, err := inGame.NewRouter(ctl)
	if err != nil {
		log.Fatalf("error on build router.: %v", err)
	}
	room.SetRoutes(e.RouterGroup)

	return nil

}
