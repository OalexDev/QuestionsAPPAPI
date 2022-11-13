package model

type Health struct {
	Message string
	Routes  []Route
}

type Route struct {
	Method  string
	Path    string
	Message string `json:"message,omitempty"`
}
