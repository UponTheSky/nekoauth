package auth

// customized errors
// this part will be moved into a common utils package in the main backend project
type HttpError struct {
	message string
	Status  int
}

func NewHttpError(message string, status int) *HttpError {
	return &HttpError{
		message: message,
		Status:  status,
	}
}

func (e *HttpError) Error() string {
	return e.message
}
