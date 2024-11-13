package user

import (
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"tempalte_model/internal/interface/entity"
	"tempalte_model/internal/interface/mock"
	"testing"
)

func Test_iUser_Register(t *testing.T) {
	gone.RunTest(func(in struct {
		iUser      *iUser               `gone:"*"` //inject iUser for test
		iDependent *mock.MockIDependent `gone:"*"` //inject iDependent for mock
	}) {
		err := gone.ToError("err")
		in.iDependent.EXPECT().DoSomething().Return(err)

		register, err2 := in.iUser.Register(&entity.RegisterParam{
			Username: "test",
			Password: "test",
		})
		assert.Nil(t, register)
		assert.Equal(t, err2, err)
	}, func(cemetery gone.Cemetery) error {
		controller := gomock.NewController(t)

		//load all mocked components
		mock.MockPriest(cemetery, controller)

		//bury the tested component
		cemetery.Bury(&iUser{})

		return nil
	})
}
