package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GolangPruebaRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIdHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s),
		getProductByIdRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s),
		getProductsRequestDecoder, kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := kithttp.NewServer(makeAddProductsEndPoint(s),
		addProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := kithttp.NewServer(makeUpdateProductEndPoint(s),
		updateProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductHandler)

	deleteProductHandler := kithttp.NewServer(makeDeleteProductEndPoint(s),
		getDeleteRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	getBestSellersHandler := kithttp.NewServer(makeBestSellerEndpoint(s),
		getBestSellerRecuestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/bestsellers", getBestSellersHandler)
	return r
}

func getProductByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIdRequest{
		ProductID: productId,
	}, nil
}

func getProductsRequestDecoder(contex context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil

}

func addProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func updateProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
func getDeleteRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return deleteProductRequest{
		ProductID: chi.URLParam(r, "id"),
	}, nil
}

func getBestSellerRecuestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getBestSellersRequest{}, nil
}
