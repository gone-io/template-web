package user

import (
	"template_module/internal/interface/entity"
	"template_module/internal/interface/service"

	"github.com/gone-io/gone"
)

type iUser struct {
	gone.Flag

	iDep service.IDependent `gone:"*"`
}

func (s *iUser) Register(registerParam *entity.RegisterParam) (*entity.LoginResult, error) {
	err := s.iDep.DoSomething()
	if err != nil {
		return nil, gone.ToError(err)
	}

	return &entity.LoginResult{
		Token: "token",
		User: &entity.User{
			Id:   1,
			Name: "Dapeng",
		},
	}, nil
}

func (s *iUser) GetUserIdFromToken(token string) (userId int64, err error) {
	return 1, nil
}

func (s *iUser) Login(loginParam *entity.LoginParam) (*entity.LoginResult, error) {
	return &entity.LoginResult{
		Token: "token",
		User: &entity.User{
			Id:   1,
			Name: "Dapeng",
		},
	}, nil
}

func (s *iUser) Logout(token string) error {
	return nil
}

func (s *iUser) GetUserById(userId int64) (*entity.User, error) {
	return &entity.User{
		Id:   1,
		Name: "Dapeng",
	}, nil
}
