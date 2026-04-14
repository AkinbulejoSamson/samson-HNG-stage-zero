package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/handler"
	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/service"
)

func TestClassifyName(t *testing.T) {
	genderService := service.NewGenderService()
	genderHanler := handler.NewGenderHandler(genderService)

	tests := []struct {
		Name               string
		QueryParam         string
		ExpectedStatusCode int
		ExpectedBody       string
	}{
		{
			Name:               "Empty Name",
			QueryParam:         "",
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedBody:       "name is required",
		},
		{
			Name:               "Numeric Name",
			QueryParam:         "12345",
			ExpectedStatusCode: http.StatusUnprocessableEntity,
			ExpectedBody:       "name is not a string",
		},
		{
			Name:               "Valid Name",
			QueryParam:         "samson",
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody:       "success",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/classify?name="+test.QueryParam, nil)

			rr := httptest.NewRecorder()
			genderHanler.ClassifyName(rr, req)
			if rr.Code != test.ExpectedStatusCode {
				t.Errorf("%s: Expected status code %d, got %d", test.Name, test.ExpectedStatusCode, rr.Code)
			}
			if !strings.Contains(rr.Body.String(), test.ExpectedBody) {
				t.Errorf("%s: Expected body %s, got %s", test.Name, test.ExpectedBody, rr.Body.String())
			}

			if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
				t.Errorf("%s: missing CORS header", test.Name)
			}
		})
	}
}
