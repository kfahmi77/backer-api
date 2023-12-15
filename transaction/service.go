package transaction

import (
	"errors"
	"go-backer-api/campaign"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
}

func NewService(repository Repository, repository2 campaign.Repository) *service {
	return &service{repository, repository2}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not the owner of owner transaction")
	}

	transactions, err := s.repository.GetCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
