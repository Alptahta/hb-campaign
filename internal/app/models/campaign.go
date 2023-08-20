package models

type Campaign struct {
	ID                     int64
	Name                   string
	ProductCode            string
	Duration               int
	PriceManipulationLimit int
	TargetSalesCount       int
	StartedAt              string
	FinishedAt             string
	Status                 string
}

type CreateCampaignRequest struct {
	Name                   string
	ProductCode            string
	Duration               int
	PriceManipulationLimit int
	TargetSalesCount       int
}

type CreateCampaignDTO struct {
	Name                   string
	ProductCode            string
	Duration               int
	PriceManipulationLimit int
	TargetSalesCount       int
	StartedAt              string
	FinishedAt             string
	Status                 string
}

type CampaignWithFinishTime struct {
	ID         int64
	Name       string
	FinishedAt string
}
