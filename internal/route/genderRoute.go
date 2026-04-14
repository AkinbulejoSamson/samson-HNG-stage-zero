package route

import (
	"net/http"

	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/handler"
	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/service"
)

func SetupGenderRoutes(genderService service.GenderService) *http.ServeMux {
	genderMux := http.NewServeMux()
	genderHandler := handler.NewGenderHandler(genderService)

	//routes
	genderMux.HandleFunc("GET /api/classify", genderHandler.ClassifyName)

	return genderMux
}
