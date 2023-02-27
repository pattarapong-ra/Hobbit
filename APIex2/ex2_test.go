package main

import (
	"reflect"
	"testing"
	"time"
)

func TestGetPromo(t *testing.T) {
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
		{
			name: "Success",
			args: args{
				currentDate: "2020-12-30",
			},
			want: promotion{
				PromoName:    "Promo3",
				InterestRate: 25,
			},
		},
		{
			name: "Success",
			args: args{
				currentDate: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC).Format("2006-01-02"),
			},
			want: promotion{
				PromoName:    "Promo2",
				InterestRate: 18,
			},
		},
		{
			name: "Fail",
			args: args{
				currentDate: "2020-01",
			},
			want: promotion{
				PromoName:    "",
				InterestRate: 0,
			},
		},
		{
			name: "Fail",
			args: args{
				currentDate: "2023-02-27",
			},
			want: promotion{
				PromoName:    "",
				InterestRate: 0,
			},
		},
		{
			name: "Fail",
			args: args{
				currentDate: "asdsadda",
			},
			want: promotion{
				PromoName:    "",
				InterestRate: 0,
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
}
