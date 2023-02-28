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
				num1: 10,
				num2: -5,
			},
			want: number{
				num1: 10,
				num2: -5,
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

func Test_number_show(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name string
		args fields
	}{
		{
			name: "Success",
			args: fields{
				num1: 50,
				num2: 50,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.args.num1,
				num2: tt.args.num2,
			}
			n.show()
		})
	}
}

func Test_number_Plus(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Success",
			fields: fields{
				num1: 50,
				num2: 50,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.Plus(); got != tt.want {
				t.Errorf("number.Plus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_Minus(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Success",
			fields: fields{
				num1: 50,
				num2: 50,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.Minus(); got != tt.want {
				t.Errorf("number.Minus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_Multi(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Success",
			fields: fields{
				num1: 50,
				num2: 50,
			},
			want: 2500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.Multi(); got != tt.want {
				t.Errorf("number.Multi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_Div(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Success",
			fields: fields{
				num1: 50,
				num2: 50,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.Div(); got != tt.want {
				t.Errorf("number.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculator(t *testing.T) {
	type args struct {
		n mathOperator
		c int64
	}
	tests := []struct {
		name string
		args args
		want number
	}{
		{
			name: "Swap Success",
			args: args{
				n: number{
					num1: 5,
					num2: 10,
				},
				c: 6,
			},
			want:number{
				num1:10,
				num2:5,
			},
		},
		{
			name: "Exit Success",
			args: args{
				n: number{
					num1: 5,
					num2: 10,
				},
				c: 7,
			},
			want:number{
				num1:5,
				num2:10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculator(tt.args.n, tt.args.c)
		})
	}
}
