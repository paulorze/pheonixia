package customErrors

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

var ServerError = CustomError {Code: 500, Message: "server error"}

var InvalidId = CustomError {Code: 422, Message: "id is a required field"}

var MockError = CustomError {Code: 420, Message: "Enhance your calm"}