package main

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(product *Product) error
	HasHistoryToday(id uint) bool
	CreateHistory(product *Product) error
	CalcRealValues(product *Product) error
}

type ProductRepository struct {
	db *gorm.DB
}

func (repo *ProductRepository) Create(product *Product) error {
	toAssign := Product{
		Name:           product.Name,
		Link:           product.Link,
		Images:         product.Images,
		ListPrice:      product.ListPrice,
		Price:          product.Price,
		PriceWithCC:    product.PriceWithCC,
		Discount:       CalcDiscount(product.ListPrice, product.Price),
		DiscountWithCC: CalcDiscount(product.ListPrice, product.Price),
		HasStock:       product.HasStock,
	}
	log.Debugf("Discount: %d - DiscountWithCC: %d", toAssign.Discount, toAssign.DiscountWithCC)
	if err := repo.db.Where(Product{Sku: product.Sku}).Assign(toAssign).FirstOrCreate(&product).Error; err != nil {
		return err
	}
	return nil
}

func (repo ProductRepository) HasHistoryToday(id uint) bool {
	query := `SELECT EXISTS(
			SELECT 1 
			FROM product_histories 
			WHERE DATE(created_at) = CURDATE() AND product_id = ? 
			LIMIT 1
		)`

	var exists bool
	err := repo.db.Raw(query, id).Row().Scan(&exists)
	if err != nil {
		log.Error(err)
		return false
	}
	return exists
}

func (repo *ProductRepository) CreateHistory(product *Product) error {
	if product.ID == 0 {
		return errors.New("the product should have an ID")
	}

	productHistory := ProductHistory{
		ListPrice:   product.ListPrice,
		Price:       product.Price,
		PriceWithCC: product.PriceWithCC,
	}
	if repo.HasHistoryToday(product.ID) {
		log.Debug("there is a record for today")
		ph := ProductHistory{}
		if err := repo.db.Model(&ph).Where("product_id = ? AND DATE(created_at) = CURDATE()", product.ID).Updates(productHistory).Error; err != nil {
			log.Error(err)
			return err
		}
	} else {
		log.Debug("creating record for today")
		productHistory.ProductID = product.ID
		if err := repo.db.Create(&productHistory).Error; err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func (repo *ProductRepository) CalcRealValues(product *Product) error {
	if product.ID == 0 {
		return errors.New("the product should have an ID")
	}

	query := `SELECT 
		(SELECT list_price
		 FROM
			 (SELECT list_price,
					 count(list_price)
			  FROM product_histories
			  WHERE created_at >= now() - interval 3 MONTH
				  AND product_id = ?
			  GROUP BY list_price
			  LIMIT 1) AS lp) AS list_price,
		(SELECT price
		 FROM
			 (SELECT price,
					 count(price)
			  FROM product_histories
			  WHERE created_at >= now() - interval 3 MONTH
				  AND product_id = ?
			  GROUP BY price
			  LIMIT 1) AS p) AS price,
		(SELECT price_with_cc
		 FROM
			 (SELECT price_with_cc,
					 count(price_with_cc)
			  FROM product_histories
			  WHERE created_at >= now() - interval 3 MONTH
				  AND product_id = ?
			  GROUP BY price_with_cc
			  LIMIT 1) AS pcc) AS price_with_cc
		`
	var listPrice uint
	var price uint
	var priceWithCC uint

	id := product.ID
	err := repo.db.Raw(query, id, id, id).Row().Scan(&listPrice, &price, &priceWithCC)
	if err != nil {
		log.Error(err)
		return err
	}

	toAssign := Product{
		ModePrice:       price,
		ModePriceWithCC: priceWithCC,
		ModeListPrice:   listPrice,
		Discount:        CalcDiscount(listPrice, product.Price),
		DiscountWithCC:  CalcDiscount(listPrice, product.PriceWithCC),
	}

	log.Debugf("Discount: %d - DiscountWithCC: %d", toAssign.Discount, toAssign.DiscountWithCC)
	if err := repo.db.Model(&product).Updates(toAssign).Error; err != nil {
		log.Errorf("can't update the real values for product ID: %d", product.ID)
		return err
	}
	log.Infof("real values updated for product ID: %d - Name: %s", product.ID, product.Name)
	return nil
}
