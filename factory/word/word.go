package word

import (
	"database/sql"
	"log"

	svcWord "github.com/Five-Series/questions/app/word"
	"github.com/Five-Series/questions/infra/environment"
	inWord "github.com/Five-Series/questions/input/word"
	outword "github.com/Five-Series/questions/output/word"
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

	// repo := outHealthCheck.NewRepository(e.DbConnection, e.RoutesGin, e.Env)

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
