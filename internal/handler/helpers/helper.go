package helpers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/dto"
)

func WriteJSONerror(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(dto.ErrorResponse{
		Status:  "error",
		Message: message,
	})
	if err != nil {
		return
	}
}

func WriteJSONSuccess(w http.ResponseWriter, statusCode int, data *dto.ProcessedData) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(dto.SuccessResponse{
		Status: "success",
		Data:   *data,
	})
	if err != nil {
		return
	}
}

func IsAlpha(name string) bool {
	return regexp.MustCompile("^[A-Za-z ]+$").MatchString(name)
}
