package service

import (
	"fmt"
	"github.com/Agilistikmal/venti/database"
	"github.com/Agilistikmal/venti/model"
)

func CreateProduct(request *model.Product) *model.Product {
	database.DB.Create(&request)
	return request
}

func UpdateProduct(productId string, request *model.Product) (*model.Product, error) {
	_, err := FindProductById(productId)
	if err != nil {
		return &model.Product{}, err
	}
	var product model.Product
	database.DB.First(&product, "id = ?", productId)

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.Stock = request.Stock
	product.Sold = request.Sold
	product.Stars = request.Stars
	product.Usage = request.Usage

	database.DB.Save(&product)
	return &product, nil
}

func DeleteProduct(productId string) error {
	_, err := FindProductById(productId)
	if err != nil {
		return err
	}
	var product model.Product
	database.DB.Where("id = ?", productId).Delete(&product)
	return nil
}

func FindAllProduct() []model.Product {
	var products []model.Product
	database.DB.Find(&products)
	return products
}

func FindProductById(productId string) (model.Product, error) {
	var product model.Product
	result := database.DB.First(&product, "id = ?", productId)
	if result.Error != nil {
		return model.Product{}, fmt.Errorf("product with id %s not found", productId)
	}
	return product, nil
}
