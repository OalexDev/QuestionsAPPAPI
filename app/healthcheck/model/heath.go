package model

type Health struct {
	Message string
	Routes  []Route
}

type Route struct {
	Method      string
	Version     string `env:"version"`
	Environment string `env:"env"`
	Path        string
	Message     string `json:"message,omitempty" `
}
