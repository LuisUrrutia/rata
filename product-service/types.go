package main

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type ProductHistory struct {
	gorm.Model
	ProductID   uint `gorm:"primary_key;auto_increment:false"`
	ListPrice   uint
	Price       uint
	PriceWithCC uint
}

type Product struct {
	gorm.Model
	Sku             string `gorm:"type:varchar(30)"`
	Name            string
	ModeListPrice   uint
	ListPrice       uint
	ModePrice       uint
	Price           uint
	ModePriceWithCC uint
	PriceWithCC     uint
	Discount        int32
	DiscountWithCC  int32
	Link            string
	Images          Images `sql:"TYPE:json"`
	StoreID         uint
	History         []ProductHistory
	HasStock        bool
	Category        string
}

type Images struct {
	Urls []string
}

func (c Images) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Images) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
