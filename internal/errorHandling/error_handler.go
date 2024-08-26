package errorHandling

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Error struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func NewError(message string, description string, status int) Error {
	return Error{
		Message:     message,
		Description: description,
		Status:      status,
	}
}

func New404Error(message string, description string) Error {
	return NewError(message, description, http.StatusNotFound)
}

func New500Error(message string, description string) Error {
	return NewError(message, description, http.StatusInternalServerError)
}

func New400Error(message string, description string) Error {
	return NewError(message, description, http.StatusBadRequest)
}

func New401Error(message string, description string) Error {
	return NewError(message, description, http.StatusUnauthorized)
}

func New403Error(message string, description string) Error {
	return NewError(message, description, http.StatusForbidden)
}

func UnAuthorizedError() Error {
	return New401Error("Unauthorized", "You are not authorized to access this resource")
}

func MissingParameterError(parameter string) Error {
	return New400Error("Missing Parameter", "The parameter "+parameter+" is required")
}

func MissingParametersError(parameters []string) Error {
	return New400Error("Missing Parameters", "The parameters "+strings.Join(parameters, ", ")+" are required")
}

func InvalidParameterError(parameter string) Error {
	return New400Error("Invalid Parameter", "The parameter "+parameter+" is invalid")
}

func InvalidParametersError(parameters []string) Error {
	return New400Error("Invalid Parameters", "The parameters "+strings.Join(parameters, ", ")+" are invalid")
}

func InternalServerError() Error {
	return New500Error("Internal Server Error", "An unexpected error occurred")
}

func DescribedInternalServerError(description string) Error {
	return New500Error("Internal Server Error", description)
}

func (e Error) Error() string {
	return e.Message
}

func SendError(w http.ResponseWriter, err Error) {
	w.WriteHeader(err.Status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
	return
}
