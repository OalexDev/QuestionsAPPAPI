package httphandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// WriteMissingParametersError  responsible for build return http on missing parameter
func WriteMissingParametersError(context *gin.Context, parameters []string) error {
	writeParametersError(context, parameters, "missing parameters")
	return nil
}

// WriteMissingHeaders  responsible for build return http on missing parameter
func WriteMissingHeadersError(context *gin.Context, parameters []string) error {
	writeParametersError(context, parameters, "missing headers")
	return nil
}

// WriteInvalidParametersError  responsible for build return http on invalid parameter
func WriteInvalidParametersError(context *gin.Context, parameters []string) error {
	writeParametersError(context, parameters, "invalid parameters")
	return nil
}

func writeParametersError(context *gin.Context, parameters []string, errorMessage string) {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: true,
			Message:  fmt.Sprintf("error: %s: %s", errorMessage, strings.Join(parameters, ", ")),
		},
	}
	context.JSON(http.StatusBadRequest, httpError)

}

// WriteBadRequest  responsible for build return http on bad request
func WriteBadRequest(context *gin.Context, message string) error {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: true,
			Message:  fmt.Sprintf("error: %s", message),
		},
	}
	context.JSON(http.StatusBadRequest, httpError)

	return nil
}

// WriteDuplicatedError responsible for build return http on duplicate errors
func WriteDuplicatedError(context *gin.Context, entityName string) error {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: true,
			Message:  fmt.Sprintf("error: the entity %s already has one inserted", entityName),
		},
	}
	context.JSON(http.StatusConflict, httpError)

	return nil
}

// WriteNotFoundSuccess responsible for build return http on not found errors
func WriteNotFoundSuccess(context *gin.Context, entityName string) error {
	httpResult := HTTPResult{
		Success: true,
		Data: HTTPGenericMessage{
			Message: fmt.Sprintf("error: no one %s was found", entityName),
		},
	}
	context.JSON(http.StatusNotFound, httpResult)

	return nil
}

// WriteGenericError responsible for build return http on generic errors
func WriteGenericError(context *gin.Context, business bool, statusCode int, message string) error {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: business,
			Message:  fmt.Sprintf("error: %s", message),
		},
	}
	context.JSON(statusCode, httpError)

	return nil
}

// WriteSuccess  responsible for build return http on generic success
func WriteSuccess(context *gin.Context, statusCode int, entity string, data interface{}) error {
	httpSuccess := HTTPResult{
		Success: true,
		Data: map[string]interface{}{
			entity: data,
		},
	}
	context.JSON(statusCode, httpSuccess)

	return nil
}

// WriteConflictError responsible for build return http on conflict errors
func WriteConflictError(context *gin.Context, entityName string) error {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: true,
			Message:  fmt.Sprintf("error: the entity %s ", entityName),
		},
	}
	context.JSON(http.StatusConflict, httpError)

	return nil
}

// WriteConflictError responsible for build return http on conflict errors
func WriteNotImplementedError(context *gin.Context) error {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: true,
			Message:  fmt.Sprintf("not been implemented"),
		},
	}
	context.JSON(http.StatusNotImplemented, httpError)

	return nil
}

// WriteInternalServerError  responsible for build return http on bad request
func WriteInternalServerError(context *gin.Context, message string) error {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: false,
			Message:  fmt.Sprintf("error: %s", message),
		},
	}
	context.JSON(http.StatusInternalServerError, httpError)

	return nil
}

// WriteUnprocessableEntity  responsible for build return http on bad request
func WriteUnprocessableEntity(context *gin.Context, message string) error {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: true,
			Message:  fmt.Sprintf("error: %s", message),
		},
	}
	context.JSON(http.StatusUnprocessableEntity, httpError)

	return nil
}

// WriteNotAllowed responsible for build return http on bad request
func WriteNotAllowed(context *gin.Context, message string) error {
	httpError := HTTPResult{
		Success: false,
		Error: &HTTPError{
			Business: true,
			Message:  fmt.Sprintf("error: %s", message),
		},
	}
	context.JSON(http.StatusForbidden, httpError)

	return nil
}
