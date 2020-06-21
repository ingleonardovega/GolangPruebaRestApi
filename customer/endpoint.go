package customer

import (
	"context"

	"github.com/GolangPruebaRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getCustomersRequest struct {
	Limit  int
	Offset int
}

// @Summary Lista de Clientes
// @Tags Customers
// @Accept json
// @Produce  json
// @Param request body customer.getCustomersRequest true "User Data"
// @Success 200 {object} customer.CustomerList "ok"
// @Router /customers/paginated [post]
func makeGetCustomersEndpoint(s Service) endpoint.Endpoint {
	getCustomersEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getCustomersRequest)
		result, err := s.GetCustomers(&req)
		helper.Catch(err)
		return result, nil
	}
	return getCustomersEndpoint
}
