package main

import "fmt"

type Product struct {
	productID       int
	productName     string
	productPrice    int
	productQuantity int
}

func ShowProduct(p []Product) {
	fmt.Println("**************************")
	fmt.Println("รายการสินค้า")
	for _, val := range p {
		if val.productQuantity == 0 {
			fmt.Printf("รหัสสินค้า : %d\nชื่อสินค้า : %s\nราคาสินค้า : %d\nจำนวนสินค้า : สินค้าหมด\n\n", val.productID, val.productName, val.productPrice)
		} else {
			fmt.Printf("รหัสสินค้า : %d\nชื่อสินค้า : %s\nราคาสินค้า : %d\nจำนวนสินค้า : %d\n\n", val.productID, val.productName, val.productPrice, val.productQuantity)
		}

	}
	fmt.Println("**************************")
}

func GetProduct(pL []Product, choice, buyQuantity, payMoney int) []Product {
	var existed bool
	if len(pL) == 0 {
		fmt.Println("ไม่มีสินค้าในรายการสินค้า กรุณาลองใหม่ภายหลัง")
		return pL
	}
	var sumPrice, change int
	var changeThousand, changeFiveHundred, changeOneHundred, changeFifty, changeTwenty, changeTen, changeFive, changeOne int

	if choice == -1 {
		return pL
	}
	for _, val := range pL {
		if choice == val.productID {
			existed = true
			break
		}
	}
	for existed != true {
		fmt.Println("ไม่มีสินค้านี้ในระบบ กรุณากรอกรหัสใหม่อีกครั้ง")
		return pL
	}

	if buyQuantity == -1 {
		return pL
	}
	if pL[choice-1].productQuantity < buyQuantity || buyQuantity == 0 {
		return pL
	}
	//if not enough product
	sumPrice = pL[choice-1].productPrice * buyQuantity
	fmt.Printf("ราคารวม : %d\n", sumPrice)

	if payMoney == -1 {
		return pL
	}

	for sumPrice > payMoney {
		return pL
	}

	pL[choice-1].productQuantity -= buyQuantity
	fmt.Println("ทำรายการสำเร็จ")

	change = payMoney - sumPrice
	if change == 0 {
		fmt.Println("ไม่มีเงินทอน")
	} else {
		fmt.Printf("เงินทอน : %d\n", change)
		changeThousand = change / 1000
		change -= changeThousand * 1000
		changeFiveHundred = change / 500
		change -= changeFiveHundred * 500
		changeOneHundred = change / 100
		change -= changeOneHundred * 100
		changeFifty = change / 50
		change -= changeFifty * 50
		changeTwenty = change / 20
		change -= changeTwenty * 20
		changeTen = change / 10
		change -= changeTen * 10
		changeFive = change / 5
		change -= changeFive * 5
		changeOne = change
		fmt.Printf("แบ๊งค์พัน : %d\n", changeThousand)
		fmt.Printf("แบ๊งค์ห้าร้อย : %d\n", changeFiveHundred)
		fmt.Printf("แบ๊งค์ร้อย : %d\n", changeOneHundred)
		fmt.Printf("แบ๊งค์ห้าสิบ : %d\n", changeFifty)
		fmt.Printf("แบ๊งค์ยี่สิบ : %d\n", changeTwenty)
		fmt.Printf("เหรียญสิบ : %d\n", changeTen)
		fmt.Printf("เหรียญห้า : %d\n", changeFive)
		fmt.Printf("เหรียญบาท : %d\n", changeOne)
	}
	ShowProduct(pL)
	return pL
}

func AddProduct(pL []Product, name string, price, quantity int) []Product {
	var product Product
	var itemIndex int
	dupe := false

	product.productID = len(pL) + 1
	product.productName = name
	product.productPrice = price
	product.productQuantity += quantity
	if name == "-1" {
		return pL
	}

	if price == -1 {
		return pL
	}

	if quantity == -1 {
		return pL
	}

	for i := 0; i < len(pL); i++ {
		if product.productName == pL[i].productName && pL[i].productPrice == product.productPrice {
			dupe = true
			itemIndex = i
		}
	}

	if dupe == true {
		pL[itemIndex].productQuantity += quantity
		ShowProduct(pL)
		return pL
	} else {
		pL = append(pL, product)
		ShowProduct(pL)
		return pL
	}
}

func main() {
	productList := []Product{}
	productList = AddProduct(productList, "llama", 20, 100)
	fmt.Println(productList)
	productList = GetProduct(productList, 1, 20, 400)
}
