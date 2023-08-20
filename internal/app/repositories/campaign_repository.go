package repositories

import (
	"database/sql"
	"fmt"
	"hb-campaign/internal/app/models"
)

type CampaignRepositoryI interface {
	CreateCampaign(cc models.CreateCampaignDTO) (int64, error)
	GetCampaignByName(campaignCode string) (models.Campaign, error)
	GetAllCampaignsWithFinishTimes() ([]models.CampaignWithFinishTime, error)
	UpdateStatusByCampaignName(campaignName, status string) error
	GetAllActiveCampaigns() ([]models.Campaign, error)
	GetAllEndedCampaigns() ([]models.Campaign, error)
}

type CampaignRepository struct {
	db *sql.DB
}

func NewCampaignRepository(db *sql.DB) *CampaignRepository {
	return &CampaignRepository{
		db: db,
	}
}

func (pr *CampaignRepository) CreateCampaign(c models.CreateCampaignDTO) (int64, error) {
	result, err := pr.db.Exec("INSERT INTO campaigns (name, product_code, duration, price_manipulation_limit, target_sales_count, started_at, finished_at, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		c.Name, c.ProductCode, c.Duration, c.PriceManipulationLimit, c.TargetSalesCount, c.StartedAt, c.FinishedAt, c.Status)
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
	if err := row.Scan(&Campaign.ID, &Campaign.Name, &Campaign.ProductCode, &Campaign.Duration, &Campaign.PriceManipulationLimit, &Campaign.TargetSalesCount, &Campaign.StartedAt, &Campaign.FinishedAt, &Campaign.Status); err != nil {
		if err == sql.ErrNoRows {
			return Campaign, fmt.Errorf("campaignByCampaignCode %s: no such Campaign", name)
		}
		return Campaign, fmt.Errorf("campaignByCampaignCode %s: %v", name, err)
	}
	return Campaign, nil
}

func (pr *CampaignRepository) GetAllCampaignsWithFinishTimes() ([]models.CampaignWithFinishTime, error) {
	var campaigns []models.CampaignWithFinishTime

	rows, err := pr.db.Query("SELECT id,name,finished_at FROM campaigns")
	if err != nil {
		return nil, fmt.Errorf("all campaigns %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var campaign models.CampaignWithFinishTime
		if err := rows.Scan(&campaign.ID, &campaign.Name, &campaign.FinishedAt); err != nil {
			return nil, fmt.Errorf("all campaigns %v", err)
		}
		campaigns = append(campaigns, campaign)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("campaigns %v", err)
	}
	return campaigns, nil
}

func (pr *CampaignRepository) UpdateStatusByCampaignName(campaignName, status string) error {
	res, err := pr.db.Exec("UPDATE campaigns SET status=? where name=?", status, campaignName)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (pr *CampaignRepository) GetAllActiveCampaigns() ([]models.Campaign, error) {
	var campaigns []models.Campaign

	rows, err := pr.db.Query("SELECT * FROM campaigns WHERE status='Active'")
	if err != nil {
		return nil, fmt.Errorf("all campaigns %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var campaign models.Campaign
		if err := rows.Scan(&campaign.ID, &campaign.Name, &campaign.ProductCode, &campaign.Duration, &campaign.Name, &campaign.TargetSalesCount, &campaign.StartedAt, &campaign.FinishedAt, &campaign.Status); err != nil {
			return nil, fmt.Errorf("all campaigns %v", err)
		}
		campaigns = append(campaigns, campaign)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("campaigns %v", err)
	}
	return campaigns, nil
}

func (pr *CampaignRepository) GetAllEndedCampaigns() ([]models.Campaign, error) {
	var campaigns []models.Campaign

	rows, err := pr.db.Query("SELECT * FROM campaigns WHERE status='Ended'")
	if err != nil {
		return nil, fmt.Errorf("all campaigns %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var campaign models.Campaign
		if err := rows.Scan(&campaign.ID, &campaign.Name, &campaign.ProductCode, &campaign.Duration, &campaign.Name, &campaign.TargetSalesCount, &campaign.StartedAt, &campaign.FinishedAt, &campaign.Status); err != nil {
			return nil, fmt.Errorf("all campaigns %v", err)
		}
		campaigns = append(campaigns, campaign)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("campaigns %v", err)
	}
	return campaigns, nil
}
