package products

import "product-api/app/utils/errors"

var (
	pro_list = []Product{}
)

func (pro *Product) Add() *errors.ResErr {
	pro_list = append(pro_list, *pro)
	return nil
}
func (pro *Product) Get() []Product {
	return pro_list
}
