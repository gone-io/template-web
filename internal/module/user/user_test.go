package user

import (
	"github.com/gone-io/gone"
	"reflect"
	"testing"
	"web/internal/interface/service"
)

func Test_iUser_Register(t *testing.T) {
	type fields struct {
		Flag gone.Flag
		iDep service.IDependent
	}
	type args struct {
		registerParam *entity.RegisterParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.LoginResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &iUser{
				Flag: tt.fields.Flag,
				iDep: tt.fields.iDep,
			}
			got, err := s.Register(tt.args.registerParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
