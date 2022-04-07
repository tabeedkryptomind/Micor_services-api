package productservice

import (
	"product-api/app/models/products"
	"product-api/app/utils/errors"
)

var counter int = 0

func AddProduct(pro products.Product) (*products.Product, *errors.ResErr) {
	if err := pro.Validate(); err != nil {
		return nil, err
	}
	pro.Add()
	return &pro, nil
}

func GetProduct() []products.Product {
	var pro products.Product
	return pro.Get()
}
