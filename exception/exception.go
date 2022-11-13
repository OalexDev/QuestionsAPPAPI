package exception

import (
	"fmt"
	"strings"
)

type NotFoundError struct {
	Entity string
}

func NewNotFoundError(entity string) NotFoundError {
	return NotFoundError{
		Entity: entity,
	}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Entity)
}

type InvalidParametersError struct {
	Entity []string
}

func NewInvalidParametersError(entity []string) *InvalidParametersError {
	return &InvalidParametersError{
		Entity: entity,
	}
}

func (ip InvalidParametersError) Error() string {
	return strings.Join(ip.Entity, ",")
}

type ConflictError struct {
	Entity string
}

func NewConflictError(entity string) ConflictError {
	return ConflictError{
		Entity: entity,
	}
}

func (c ConflictError) Error() string {
	return fmt.Sprintln(c.Entity)
}

type AuthRequired struct {
	Entity string `json:"entity"`
}

func NewAuthRequiredError(entity string) AuthRequired {
	return AuthRequired{
		Entity: entity,
	}
}

func (c AuthRequired) Error() string {
	return fmt.Sprintf("Login error on %s", c.Entity)
}

type InternalServerError struct {
	Entity string `json:"entity"`
}

func NewInternalServerError(entity string) InternalServerError {
	return InternalServerError{
		Entity: entity,
	}
}

func (c InternalServerError) Error() string {
	return fmt.Sprintf("Login error on %s", c.Entity)
}

type BadRequest struct {
	Entity string `json:"entity"`
}

func NewBadRequestError(entity string) BadRequest {
	return BadRequest{
		Entity: entity,
	}
}

func (c BadRequest) Error() string {
	return fmt.Sprintf("Login error on %s", c.Entity)
}

type Unauthorized struct {
	Entity string `json:"entity"`
}

func NewUnauthorizedError(entity string) Unauthorized {
	return Unauthorized{
		Entity: entity,
	}
}

func (c Unauthorized) Error() string {
	return fmt.Sprintf("Login error on %s", c.Entity)
}

type Forbidden struct {
	Entity string `json:"entity"`
}

func NewForbiddenError(entity string) Forbidden {
	return Forbidden{
		Entity: entity,
	}
}

func (c Forbidden) Error() string {
	return fmt.Sprintf("Login error on %s", c.Entity)
}
