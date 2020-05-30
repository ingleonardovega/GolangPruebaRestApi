package employee

import (
	"context"

	"github.com/GolangPruebaRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getEmployeesResquest struct {
	Limit  int
	Offset int
}

func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {
	getEmployeesEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getEmployeesResquest)
		result, err := s.GetEmployees(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeesEndpoint
}
