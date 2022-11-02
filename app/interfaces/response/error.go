package response

import (
	"backend/app/packages/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ValidateError struct {
	Code            int                 `json:"code"`
	Message         string              `json:"message"`
	ValidationError []utils.ValidateErr `json:"validationError"`
}

func ReturnErrorResponse(w http.ResponseWriter, code int, msg string) error {
	body := &Error{
		Code:    code,
		Message: msg,
	}
	resBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(code)
	w.Write(resBody)
	return nil
}

func ReturnValidationErrorResponse(w http.ResponseWriter, code int, msg string, validations []utils.ValidateErr) error {
	body := &ValidateError{
		Code:            code,
		Message:         msg,
		ValidationError: validations,
	}
	resBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(code)
	w.Write(resBody)
	return nil
}
