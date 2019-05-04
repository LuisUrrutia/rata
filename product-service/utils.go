package main

import (
	pb "github.com/LuisUrrutia/rata/product-service/proto/product"
)

func FromProtobufToProduct(product *pb.ProductInput) *Product {
	images := Images{Urls: product.Images}

	return &Product{
		Sku:         product.Sku,
		Name:        product.Name,
		ListPrice:   uint(product.ListPrice),
		Price:       uint(product.Price),
		PriceWithCC: uint(product.PriceWithCC),
		Link:        product.Link,
		Images:      images,
		StoreID:     uint(product.StoreId),
		HasStock:    product.HasStock,
		Category:    product.Category,
	}
}

func CalcDiscount(prev uint, actual uint) int32 {
	if actual == 0 {
		return 0
	}
	return int32(100 - (actual * 100 / prev))
}
