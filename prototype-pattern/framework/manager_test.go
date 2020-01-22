package framework_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/prototype-pattern/framework"
)

type mockProduct struct {
	createCloned bool
	framework.Product
}

func (m *mockProduct) CreateClone() framework.Product {
	return &mockProduct{
		createCloned: true,
	}
}

func TestManager_Create(t *testing.T) {
	type args struct {
		protoName string
		product   framework.Product
	}
	tests := []struct {
		name string
		args args
		want framework.Product
	}{
		{"Success", args{protoName: "hoge", product: &mockProduct{}}, &mockProduct{createCloned: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := framework.NewManager()
			m.Register(tt.args.protoName, tt.args.product)
			if got := m.Create(tt.args.protoName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
