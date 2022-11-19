package word

import (
	"database/sql"
	"log"

	svcWord "github.com/OalexDev/QuestionsAPPAPI/app/word"
	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"
	inWord "github.com/OalexDev/QuestionsAPPAPI/input/word"
	outword "github.com/OalexDev/QuestionsAPPAPI/output/word"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

type Word struct {
	DbConnection *sql.DB

	RouterGroup *gin.RouterGroup
	Env         *environment.Environment
	AWSSess     *session.Session
}

func (e *Word) Start() error {

	repo := outword.NewRepository(
		e.Env,
		e.DbConnection,
	)

	svc := svcWord.NewService(e.Env, e.AWSSess, repo)

	ctl := inWord.NewControllerWord(svc)

	word, err := inWord.NewRouter(ctl)
	if err != nil {
		log.Fatalf("error on build router.: %v", err)
	}
	word.SetRoutes(e.RouterGroup)

	return nil

}
