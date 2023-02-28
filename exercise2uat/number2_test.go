package main

import (
	"reflect"
	"testing"
)

func TestGetProduct(t *testing.T) {
	type args struct {
		pL          []Product
		choice      int
		buyQuantity int
		payMoney    int
	}
	tests := []struct {
		name string
		args args
		want []Product
	}{
		{
			name: "NoItemInList",
			args: args{
				pL:          []Product{},
				choice:      1,
				buyQuantity: 23,
				payMoney:    100,
			},
			want: []Product{},
		},
		{
			name: "SuccessNoChange",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				choice:      1,
				buyQuantity: 1,
				payMoney:    100,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 99,
					productPrice:    100},
			},
		},
		{
			name: "SuccessWithChange",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				choice:      1,
				buyQuantity: 1,
				payMoney:    500,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 99,
					productPrice:    100},
			},
		},
		{
			name: "FailedChoice",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				choice:      -1,
				buyQuantity: 10,
				payMoney:    100,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 100,
					productPrice:    100},
			},
		},
		{
			name: "FailedNoItemInList",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				choice:      322,
				buyQuantity: 10,
				payMoney:    100,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 100,
					productPrice:    100},
			},
		},
		{
			name: "FailedCancelAtQuantity",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				choice:      1,
				buyQuantity: -1,
				payMoney:    100,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 100,
					productPrice:    100},
			},
		},
		{
			name: "FailedInvalidQuantity",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				choice:      1,
				buyQuantity: 110400,
				payMoney:    100,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 100,
					productPrice:    100},
			},
		},
		{
			name: "FailedCancelMoney",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				choice:      1,
				buyQuantity: 1,
				payMoney:    -1,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 100,
					productPrice:    100},
			},
		},
		{
			name: "FailedNotEnoughMoney",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				choice:      1,
				buyQuantity: 10,
				payMoney:    100,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 100,
					productPrice:    100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProduct(tt.args.pL, tt.args.choice, tt.args.buyQuantity, tt.args.payMoney); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddProduct(t *testing.T) {
	type args struct {
		pL       []Product
		name     string
		price    int
		quantity int
	}
	tests := []struct {
		name string
		args args
		want []Product
	}{
		{
			name: "AddNewItem",
			args: args{
				pL:       []Product{},
				name:     "llama",
				price:    100,
				quantity: 100,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 100,
					productPrice:    100},
			},
		},
		{
			name: "CancelAtNameItem",
			args: args{
				pL:       []Product{},
				name:     "-1",
				price:    100,
				quantity: 100,
			},
			want: []Product{},
		},
		{
			name: "CancelAtPriceItem",
			args: args{
				pL:       []Product{},
				name:     "llama",
				price:    -1,
				quantity: 100,
			},
			want: []Product{},
		},
		{
			name: "CancelAtQuantityItem",
			args: args{
				pL:       []Product{},
				name:     "llama",
				price:    100,
				quantity: -1,
			},
			want: []Product{},
		},
		{
			name: "AddDuplicatedItem",
			args: args{
				pL: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 100,
						productPrice:    100},
				},
				name:     "llama",
				price:    100,
				quantity: 100,
			},
			want: []Product{
				{productID: 1,
					productName:     "llama",
					productQuantity: 200,
					productPrice:    100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddProduct(tt.args.pL, tt.args.name, tt.args.price, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Done",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestShowProduct(t *testing.T) {
	type args struct {
		p []Product
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "AddDuplicatedItem",
			args: args{
				p: []Product{
					{productID: 1,
						productName:     "llama",
						productQuantity: 0,
						productPrice:    100},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowProduct(tt.args.p)
		})
	}
}
