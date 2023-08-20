package services

import (
	"errors"
	"hb-campaign/internal/app/models"
	"hb-campaign/internal/app/repositories"
	"hb-campaign/internal/faketime"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

var (
	createdCampaignID int64 = 1

	mockDataCreateCampign = models.CreateCampaignRequest{
		Name:                   "C1",
		ProductCode:            "P1",
		Duration:               1,
		PriceManipulationLimit: 1,
		TargetSalesCount:       1,
	}

	mockCampaignName = "C1"

	mockCampaign = models.Campaign{
		ID:                     1,
		Name:                   "C1",
		ProductCode:            "P1",
		Duration:               1,
		PriceManipulationLimit: 1,
		TargetSalesCount:       1,
		StartedAt:              "01:00",
		FinishedAt:             "02:00",
		Status:                 "Active",
	}
)

func Test_CreateCampaign(t *testing.T) {
	t.Run("It should return an error when CreateCampaign function from repository layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockCampaignRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		cs := NewCampaignService(mockFakeTime, mockRepository)

		mockFakeTime.EXPECT().GetTime().Return("01:00").Times(1)

		mockRepository.EXPECT().
			CreateCampaign(gomock.Any()).
			Return(createdCampaignID, errors.New("fake error")).
			Times(1)

		err := cs.CreateCampaign(mockDataCreateCampign)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
	})

	t.Run("It should not return any error when CreateCampaign function from repository layer does not returns any error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockCampaignRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		cs := NewCampaignService(mockFakeTime, mockRepository)

		mockFakeTime.EXPECT().GetTime().Return("01:00").Times(1)
		mockRepository.EXPECT().
			CreateCampaign(gomock.Any()).
			Return(createdCampaignID, nil).
			Times(1)

		err := cs.CreateCampaign(mockDataCreateCampign)
		if err != nil {
			t.Fatalf("did not expect any error but got %v", err)
		}
	})
}

func Test_GetCampaignByName(t *testing.T) {
	t.Run("It should return an error when GetCampaignByName function from repository layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockCampaignRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		cs := NewCampaignService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			GetCampaignByName(gomock.Any()).
			Return(mockCampaign, errors.New("fake error")).
			Times(1)

		campaign, err := cs.GetCampaignByName(mockCampaignName)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
		if campaign != nil {
			t.Fatalf("expected nil but got %v", campaign)
		}
	})

	t.Run("It should not return any error when GetCampaignByName function from repository layer does not returns any error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockCampaignRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		cs := NewCampaignService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			GetCampaignByName(gomock.Any()).
			Return(mockCampaign, nil).
			Times(1)

		campaign, err := cs.GetCampaignByName(mockCampaignName)
		if err != nil {
			t.Fatalf("did not expect any error but got %v", err)
		}
		if campaign == nil {
			t.Fatal("expected non nil but got nil")
		}
	})
}
