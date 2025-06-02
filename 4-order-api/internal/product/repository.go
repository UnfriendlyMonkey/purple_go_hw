package product

import (
	"go/hw/4-order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (r *ProductRepository) CreateProduct(product *Product) (*Product, error) {
	err := r.Database.DB.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetProductById(id int) (*Product, error) {
	var product Product
	err := r.Database.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(product *Product) (*Product, error) {
	err := r.Database.DB.Clauses(clause.Returning{}).Updates(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	err := r.Database.DB.Delete(&Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
