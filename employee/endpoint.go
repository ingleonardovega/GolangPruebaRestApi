package employee

import (
	"context"

	"github.com/GolangPruebaRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}

type getEmployeeByIDRequest struct {
	EmployeeID string
}

type getBestEmployeeRequest struct{}

func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {
	getEmployeesEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getEmployeesRequest)
		result, err := s.GetEmployees(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeesEndpoint
}

func makeEmployeeByIDEndpoint(s Service) endpoint.Endpoint {
	getEmployeeByIDRequest := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getEmployeeByIDRequest)
		result, err := s.GetEmployeeById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeeByIDRequest
}

func makeGetBestEmployeeEndpoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndpoint := func(ct_ context.Context, request interface{}) (response interface{}, err error) {
		result, err := s.GetBestEmployee()
		helper.Catch(err)
		return result, nil
	}
	return getBestEmployeeEndpoint

}
