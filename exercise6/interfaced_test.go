package main

import (
	"reflect"
	"testing"
)

func Test_number_init(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name string
		args fields
		want number
	}{
		{
			name: "Success",
			args: fields{
				num1: 2,
				num2: 1,
			},
			want: number{
				num1: 10,
				num2: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.args.num1,
				num2: tt.args.num2,
			}
			if got := n.init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("number.init() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*func TestGetPromo(t *testing.T) {
	type args struct {
		currentDate string
	}
	tests := []struct {
		name string
		args args
		want promotion
	}{
		{
			name: "Success",
			args: args{
				currentDate: "2020-01-01",
			},
			want: promotion{
				PromoName:    "Promo1",
				InterestRate: 2.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPromo(tt.args.currentDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPromo() = %v, want %v", got, tt.want)
			}
		})
	}
}*/
