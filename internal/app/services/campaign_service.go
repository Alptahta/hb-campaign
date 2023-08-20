package services

import (
	"hb-campaign/internal/app/models"
	"hb-campaign/internal/app/repositories"
	"hb-campaign/internal/faketime"
	"strconv"
	"strings"
)

type CampaignServiceI interface {
	CreateCampaign(cp models.CreateCampaignRequest) error
	GetCampaignByName(CampaignCode string) (*models.Campaign, error)
	GetAllActiveCampaigns() ([]models.Campaign, error)
	GetAllEndedCampaigns() ([]models.Campaign, error)
	Update() error
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

func (cs CampaignService) CreateCampaign(cp models.CreateCampaignRequest) error {
	time := cs.t.GetTime()
	ss := strings.Split(time, ":")
	hour, err := strconv.Atoi(ss[0])
	if err != nil {
		return err
	}

	campaignDTO := models.CreateCampaignDTO{
		Name:                   cp.Name,
		ProductCode:            cp.ProductCode,
		Duration:               cp.Duration,
		PriceManipulationLimit: cp.PriceManipulationLimit,
		TargetSalesCount:       cp.TargetSalesCount,
		StartedAt:              time,
		FinishedAt:             strconv.Itoa(hour + cp.Duration),
		Status:                 "Active",
	}

	_, err = cs.cr.CreateCampaign(campaignDTO)
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

func (cs CampaignService) GetAllActiveCampaigns() ([]models.Campaign, error) {
	campaigns, err := cs.cr.GetAllActiveCampaigns()
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (cs CampaignService) GetAllEndedCampaigns() ([]models.Campaign, error) {
	campaigns, err := cs.cr.GetAllEndedCampaigns()
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (cs CampaignService) Update() error {
	time := cs.t.GetTime()
	ss := strings.Split(time, ":")
	currentHour, err := strconv.Atoi(ss[0])
	if err != nil {
		return err
	}

	campaigns, err := cs.cr.GetAllCampaignsWithFinishTimes()
	if err != nil {
		return err
	}

	for _, campaign := range campaigns {
		ss := strings.Split(campaign.FinishedAt, ":")
		campaignFinishTime, err := strconv.Atoi(ss[0])
		if err != nil {
			return err
		}

		if currentHour > campaignFinishTime {
			err = cs.cr.UpdateStatusByCampaignName(campaign.Name, "Ended")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
