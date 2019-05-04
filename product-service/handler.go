package main

import (
	"context"
	pb "github.com/LuisUrrutia/rata/product-service/proto/product"
	log "github.com/Sirupsen/logrus"
)

type Service struct {
	repo Repository
}

func (srv *Service) CreateOrUpdate(ctx context.Context, req *pb.ProductInput, res *pb.SuccessOutput) error {
	product := FromProtobufToProduct(req)

	err := srv.repo.Create(product)
	if err != nil {
		log.Error(err)
		res.Success = false
		res.Message = err.Error()
	} else {
		log.Printf("product saved ID: %d - SKU: %s - Name: %s", product.ID, product.Sku, product.Name)
		if err := srv.repo.CreateHistory(product); err != nil {
			log.Print("can't create history")
			log.Error(err)
		}
		if err := srv.repo.CalcRealValues(product); err != nil {
			log.Error(err)
		}
		res.Success = true
	}
	return nil
}
