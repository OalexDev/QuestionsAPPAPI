package main

import (
	"context"
	"database/sql"

	"github.com/OalexDev/QuestionsAPPAPI/factory/game"
	"github.com/OalexDev/QuestionsAPPAPI/factory/healthcheck"
	"github.com/OalexDev/QuestionsAPPAPI/factory/room"
	"github.com/OalexDev/QuestionsAPPAPI/factory/word"
	"github.com/OalexDev/QuestionsAPPAPI/infra/database"
	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	env          *environment.Environment
	ginLambda    *ginadapter.GinLambda
	app          *gin.Engine
	dbConnection *sql.DB
	routerGroup  *gin.RouterGroup
	sess         *session.Session
)

func init() {

	// DB
	env = environment.LoadOrDie()
	db := database.New(env)
	dbConnection = db.Connect()

	app = gin.Default()
	ginLambda = ginadapter.New(app)
	app.Use(corsConfig())
	routerGroup = app.Group("/v1")

	build()

	//aws
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

}

func main() {

	// Como Matar as salas?
	// Logins por Nome
	// Refactory do codig
	// Refactory dos erros
	// Logs

	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func corsConfig() gin.HandlerFunc {

	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-ROOM-ID", "X-PLAYER-ID", "Authorization"}

	handler := cors.New(configCors)
	return handler
}

func build() {

	// Game
	game := game.Game{
		DbConnection: dbConnection,
		Env:          env,
		RouterGroup:  routerGroup,
		AWSSess:      sess,
	}
	game.Start()

	// Rooms
	room := room.Room{
		DbConnection: dbConnection,
		Env:          env,
		RouterGroup:  routerGroup,
		AWSSess:      sess,
	}
	room.Start()

	// Word
	word := word.Word{
		DbConnection: dbConnection,
		Env:          env,
		RouterGroup:  routerGroup,
		AWSSess:      sess,
	}
	word.Start()

	// Health check
	healthStarter := healthcheck.Healthcheck{
		Env:          env,
		DbConnection: dbConnection,
		RouterGroup:  &app.RouterGroup,
		RoutesGin:    app.Routes(),
	}
	healthStarter.Start()

}
