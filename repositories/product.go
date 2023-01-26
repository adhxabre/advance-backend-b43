// Create package repositories here ...
package repositories

// Import the required packages here ...
import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// Declare ProductRepository interface here ...
type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

// Create RepositoryProduct function here ...
func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

// Create FindProducts method here ...
func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Find(&products).Error

	return products, err
}

// Create GetProduct method here ...
func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

// Create CreateProduct method here ...
func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}
