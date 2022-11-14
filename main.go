package main

import (
	"context"
	"database/sql"

	"github.com/Five-Series/questions/factory/healthcheck"
	"github.com/Five-Series/questions/factory/room"
	"github.com/Five-Series/questions/factory/word"
	"github.com/Five-Series/questions/infra/database"
	"github.com/Five-Series/questions/infra/environment"

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

	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func corsConfig() gin.HandlerFunc {

	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-TRACE-ID", "X-CLIENT-ID", "X-USER-TYPE", "X-USER-WMSID", "X-USER-ROLE", "X-USER-NAME", "Authorization"}

	handler := cors.New(configCors)
	return handler
}

func build() {

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
