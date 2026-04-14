package handler

import (
	"net/http"

	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/handler/helpers"
	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/service"
)

type GenderHandler struct {
	genderService service.GenderService
}

func NewGenderHandler(genderService service.GenderService) *GenderHandler {
	return &GenderHandler{genderService: genderService}
}

func (h *GenderHandler) ClassifyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	name := r.URL.Query().Get("name")

	if name == "" {
		helpers.WriteJSONerror(w, http.StatusBadRequest, "name is required")
		return
	}

	if !helpers.IsAlpha(name) {
		helpers.WriteJSONerror(w, http.StatusUnprocessableEntity, "name is not a string")
		return
	}

	data, statusCode, err := h.genderService.ClassifyName(r.Context(), name)
	if err != nil {
		if err.Error() == "no prediction available" {
			helpers.WriteJSONerror(w, statusCode, "no prediction available for this name")
			return
		}

		helpers.WriteJSONerror(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSONSuccess(w, statusCode, data)
}
