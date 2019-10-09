package util

import "github.com/Software-Architecture-2019-2/phets-event-ms/model"

// GenerateErrorResponse with the message and details of error
func GenerateErrorResponse(code int, message string, err error) model.Error {
	return model.Error{
		Code:    code,
		Message: message,
		Description: model.ErrorDescription{
			Error: err.Error(),
		},
	}
}
