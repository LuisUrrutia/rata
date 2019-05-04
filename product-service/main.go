package main

import (
	pb "github.com/LuisUrrutia/rata/product-service/proto/product"
	log "github.com/Sirupsen/logrus"
	"github.com/micro/go-micro"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ProductHistory{})
	repo := &ProductRepository{db}

	srv := micro.NewService(
		micro.Name("me.urrutia.scraper.product"),
	)
	srv.Init()

	err = pb.RegisterProductServiceHandler(srv.Server(), &Service{repo})
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
