package services

import (
	"hb-campaign/internal/app/models"
	"hb-campaign/internal/app/repositories"
	"hb-campaign/internal/faketime"
)

type CampaignServiceI interface {
	CreateCampaign(cp models.CreateCampaign) error
	GetCampaignByName(CampaignCode string) (*models.Campaign, error)
}

type CampaignService struct {
	t  faketime.TimeInterface
	cr repositories.CampaignRepositoryI
}

func NewCampaignService(t faketime.TimeInterface, cr repositories.CampaignRepositoryI) *CampaignService {
	return &CampaignService{
		t:  t,
		cr: cr,
	}
}

func (cs CampaignService) CreateCampaign(cp models.CreateCampaign) error {
	_, err := cs.cr.CreateCampaign(cp)
	if err != nil {
		return err
	}
	return nil
}

func (cs CampaignService) GetCampaignByName(CampaignCode string) (*models.Campaign, error) {
	Campaign, err := cs.cr.GetCampaignByName(CampaignCode)
	if err != nil {
		return nil, err
	}
	return &Campaign, nil
}
