// Create package repositories here ...
package repositories

// Import the required packages here ...
import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// Declare ProfileRepository interface here ...
type ProfileRepository interface {
	GetProfile(ID int) (models.Profile, error)
}

// Create RepositoryProfile function here ...
func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

// Create GetProfile method here ...
func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error

	return profile, err
}
