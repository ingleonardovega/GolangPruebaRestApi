package employee

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GolangPruebaRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getEmployeesHandler := kithttp.NewServer(makeGetEmployeesEndpoint(s),
		getEmployeesRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	return r
}

func getEmployeesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesResquest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
