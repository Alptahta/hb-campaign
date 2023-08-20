package repositories

import (
	"database/sql"
	"fmt"
	"hb-campaign/internal/app/models"
)

type CampaignRepositoryI interface {
	CreateCampaign(cc models.CreateCampaign) (int64, error)
	GetCampaignByName(campaignCode string) (models.Campaign, error)
}

type CampaignRepository struct {
	db *sql.DB
}

func NewCampaignRepository(db *sql.DB) *CampaignRepository {
	return &CampaignRepository{
		db: db,
	}
}

func (pr *CampaignRepository) CreateCampaign(c models.CreateCampaign) (int64, error) {
	result, err := pr.db.Exec("INSERT INTO campaigns (name, product_code, duration, price_manipulation_limit, target_sales_count) VALUES (?, ?, ?, ?, ?)",
		c.Name, c.ProductCode, c.Duration, c.PriceManipulationLimit, c.TargetSalesCount)
	if err != nil {
		return 0, fmt.Errorf("addCampaign: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addCampaign: %v", err)
	}
	return id, nil
}

func (pr *CampaignRepository) GetCampaignByName(name string) (models.Campaign, error) {
	var Campaign models.Campaign
	row := pr.db.QueryRow("SELECT * FROM campaigns WHERE name = ?", name)
	if err := row.Scan(&Campaign.ID, &Campaign.Name, &Campaign.ProductCode, &Campaign.Duration, &Campaign.PriceManipulationLimit, &Campaign.TargetSalesCount); err != nil {
		if err == sql.ErrNoRows {
			return Campaign, fmt.Errorf("campaignByCampaignCode %s: no such Campaign", name)
		}
		return Campaign, fmt.Errorf("campaignByCampaignCode %s: %v", name, err)
	}
	return Campaign, nil
}
