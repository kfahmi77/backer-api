package campaign

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaign(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
