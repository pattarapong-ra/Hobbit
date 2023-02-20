package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Process(t *testing.T) {
	var p []Product
	var pr Product
	pr.productID=1
	pr.productName="s"
	pr.productPrice=11
	pr.productQuantity=23
	p=append(p, pr)
	ShowProduct(p)
	assert.Equal(t, "A", "B")
}
