package input

import "github.com/OalexDev/QuestionsAPPAPI/infra/environment"

type DefaultController struct {
	Environment *environment.Environment
	TraceID     string
}

func NewDefaultController(env *environment.Environment) *DefaultController {

	return &DefaultController{
		Environment: env}

}
