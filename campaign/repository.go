package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(id int) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	if err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	if err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(id int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("id = ?", id).Preload("CampaignImages").Preload("User").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
