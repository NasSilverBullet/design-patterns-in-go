package idcard_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/factory-method-pattern/framework"
	"github.com/NasSilverBullet/design-patterns-in-go/factory-method-pattern/idcard"
)

func TestIDCardFactory_CreateProduct(t *testing.T) {
	type args struct {
		o string
	}
	tests := []struct {
		name string
		args args
		want framework.Product
	}{
		{"Success", args{o: "hoge"}, idcard.New("hoge")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icf := idcard.NewFactory()
			if got := icf.CreateProduct(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDCardFactory.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDCardFactory_RegisterProduct(t *testing.T) {
	type args struct {
		p framework.Product
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Success", args{p: idcard.New("hoge")}, []string{"hoge"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icf := idcard.NewFactory()
			icf.RegisterProduct(tt.args.p)
			if got := icf.(*idcard.IDCardFactory).GetOwners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDCardFactory.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
