package framework_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/factory-method-pattern/framework"
)

type mockProduct struct {
	owner      string
	registered bool
}

func (m *mockProduct) Use() {}

type mockFactory struct{}

func (m *mockFactory) CreateProduct(owner string) framework.Product {
	return &mockProduct{
		owner: owner,
	}
}

func (m *mockFactory) RegisterProduct(product framework.Product) {
	product.(*mockProduct).registered = true
}

func Test_factory_Create(t *testing.T) {
	type args struct {
		owner string
	}
	tests := []struct {
		name string
		args args
		want framework.Product
	}{
		{"Success", args{owner: "hoge"}, &mockProduct{owner: "hoge", registered: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := framework.NewFactory(&mockFactory{})
			if got := f.Create(tt.args.owner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factory.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
