package product

import (
	"context"

	"github.com/GolangPruebaRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getProductByIdRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type updateProductRequest struct {
	ID           int64
	Category     string
	Description  string
	ListPrice    float64
	StandardCost float64
	ProductCode  string
	ProductName  string
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductID string
}

type getBestSellersRequest struct{}

// @Summary Producto by Id
// @Tags Producto
// @Accept json
// @Produce  json
// @Param id path int true "Producto Id"
// @Success 200 {object} product.Product "ok"
// @Router /products/{id} [get]

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIdRequest)
		product, err := s.GetProductById(&req)
		helper.Catch(err)
		return product, nil
	}

	return getProductByIdEndpoint
}

// @Summary Lista de Productos
// @Tags Producto
// @Accept json
// @Produce  json
// @Param request body product.getProductsRequest true "User Data"
// @Success 200 {object} product.ProductList "ok"
// @Router /products/paginated [post]
func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		helper.Catch(err)
		return result, nil
	}
	return getProductsEndPoint
}

// @Summary Insertar Productos
// @Tags Producto
// @Accept json
// @Produce  json
// @Param request body product.getAddProductRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /products/ [post]
func makeAddProductsEndPoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		producId, err := s.InsertProduct(&req)
		helper.Catch(err)
		return producId, nil
	}
	return addProductEndpoint
}

// @Summary Update Producto
// @Tags Producto
// @Accept json
// @Produce  json
// @Param request body product.updateProductRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /products/ [put]
func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	updateProductEnpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		r, err := s.UpdateProduct(&req)
		helper.Catch(err)
		return r, nil
	}
	return updateProductEnpoint
}

// @Summary Eliminar Producto
// @Tags Producto
// @Accept json
// @Produce  json
// @Param id path int true "Producto Id"
// @Success 200 {integer} int "ok"
// @Router /products/{id} [delete]
func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	deleteProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		result, err := s.DeleteProduct(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteProductEndPoint
}

// @Summary Best Sellers
// @Tags Producto
// @Accept json
// @Produce  json
// @Success 200 {object} product.ProductTopResponse "ok"
// @Router /products/bestSellers [get]
func makeBestSellerEndpoint(s Service) endpoint.Endpoint {
	getBestSellersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestSellers()
		helper.Catch(err)
		return result, nil
	}
	return getBestSellersEndpoint

}
