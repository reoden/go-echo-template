package dtos

import "github.com/reoden/go-echo-template/internal/catalogs/products/dtos"

type GetProductByIdQueryResponse struct {
	Product *dtos.ProductDto `json:"product"`
}
