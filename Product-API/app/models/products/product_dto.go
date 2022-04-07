package products

import (
	"product-api/app/utils/errors"
	"strings"
	"time"
)

type Product struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
}

func (pro *Product) Validate() *errors.ResErr {
	pro.Name = strings.TrimSpace(pro.Name)
	pro.Title = strings.TrimSpace(pro.Title)
	pro.Created_at = time.Now()
	if pro.Title == "" {
		return errors.BadRequestError("empty product name")
	}
	if pro.Title == "" {
		return errors.BadRequestError("no dircription of of the product")
	}
	return nil

}
