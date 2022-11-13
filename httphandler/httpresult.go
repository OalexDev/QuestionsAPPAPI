package httphandler

type HTTPResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *HTTPError  `json:"error,omitempty"`
}

type HTTPError struct {
	Business bool   `json:"business,omitempty"`
	Message  string `json:"message,omitempty"`
}

type HTTPGenericMessage struct {
	Message string `json:"message,omitempty"`
}
