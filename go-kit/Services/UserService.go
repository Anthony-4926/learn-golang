package Services

type IUserService interface {
	GetName(userId uint) string
	DelUser(userId uint) error
}

type UserService struct {
}

func (this UserService) GetName(userId uint) string {
	if userId == 101 {
		return "Anthony_4926"
	}
	return "guest"
}

func (this UserService) DelUser(userId uint) error {
	return nil
}
