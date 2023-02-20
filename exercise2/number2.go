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

func main() {
	productList := []Product{}
	menuChoice := ""
	var existed bool

	ShowProduct(productList)
	for menuChoice != "exit" {
		fmt.Println("\tget")
		fmt.Println("\tadd")
		fmt.Println("\texit")
		fmt.Printf("ตัวเลือก : ")
		fmt.Scanf("%s", &menuChoice)
		switch menuChoice {
		case "get":
			if len(productList) == 0 {
				fmt.Println("ไม่มีสินค้าในรายการสินค้า กรุณาลองใหม่ภายหลัง")
				break
			}
			var productChoice, buyQuantity, sumPrice, payMoney, change int
			var changeThousand, changeFiveHundred, changeOneHundred, changeFifty, changeTwenty, changeTen, changeFive, changeOne int

			fmt.Printf("เลือกชนิดสินค้า (พิมพ์ ID)(กด -1 เพื่อยกเลิกการซื้อ) :")
			fmt.Scanf("%d", &productChoice)
			if productChoice == -1 {
				break
			}
			for _, val := range productList {
				if productChoice == val.productID {
					existed = true
					break
				}
			}
			for existed != true {
				fmt.Println("ไม่มีสินค้านี้ในระบบ กรุณากรอกรหัสใหม่อีกครั้ง")
				fmt.Printf("เลือกชนิดสินค้า (พิมพ์ ID) :")
				fmt.Scanf("%d", &productChoice)
				for _, val := range productList {
					if productChoice == val.productID {
						existed = true
					}
				}
				if existed == true {
					break
				}
			}

			fmt.Printf("จำนวนสินค้าที่ซื้อ (กด -1 เพื่อยกเลิกการซื้อ) :")
			fmt.Scanf("%d", &buyQuantity)
			if buyQuantity == -1 {
				break
			}
			for productList[productChoice-1].productQuantity < buyQuantity || buyQuantity == 0 {
				if buyQuantity == -1 {
					break
				}
				fmt.Println("จำนวนสินค้าไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง")
				fmt.Printf("จำนวนสินค้าที่ซื้อ :")
				fmt.Scanf("%d", &buyQuantity)
			}
			//if not enough product
			sumPrice = productList[productChoice-1].productPrice * buyQuantity
			fmt.Printf("ราคารวม : %d\n", sumPrice)
			fmt.Printf("จำนวนเงินที่จ่าย (กด -1 เพื่อยกเลิกการซื้อ):")
			fmt.Scanf("%d", &payMoney)
			if payMoney == -1 {
				break
			}
			if sumPrice > payMoney {
				fmt.Println("จำนวนเงินที่จ่ายไม่เพียงพอ ทำรายการไม่สำเร็จ")
				fmt.Println("กรุณาทำรายการใหม่อีกครั้ง")
			} else {
				productList[productChoice-1].productQuantity -= buyQuantity
				fmt.Println("ทำรายการสำเร็จ")
			}
			//if enough money proceed
			//if not prompt smth
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
			ShowProduct(productList)
			//print ทอนกี่แบ๊ง กี่เหรียญ
		case "add":

			var product Product
			var addProductName string
			var addProductPrice, addProductQuantity, addProductID int
			var itemIndex int
			dupe := false

			addProductID = len(productList) + 1
			product.productID = addProductID
			fmt.Printf("ชื่อของสินค้าที่เพิ่ม (กด -1 เพื่อยกเลิกการซื้อ): ")
			fmt.Scanf("%s", &addProductName)
			if addProductName == "-1" {
				break
			}
			product.productName = addProductName
			fmt.Printf("ราคาของสินค้าที่เพิ่ม (กด -1 เพื่อยกเลิกการซื้อ): ")
			fmt.Scanf("%d", &addProductPrice)
			if addProductPrice == -1 {
				break
			}
			product.productPrice = addProductPrice
			fmt.Printf("จำนวนสินค้าที่เพิ่ม (กด -1 เพื่อยกเลิกการซื้อ): ")
			fmt.Scanf("%d", &addProductQuantity)
			if addProductQuantity == -1 {
				break
			}
			product.productQuantity += addProductQuantity
			for i := 0; i < len(productList); i++ {
				if product.productName == productList[i].productName && productList[i].productPrice == product.productPrice {
					dupe = true
					itemIndex = i
				}
			}
			if dupe == true {
				productList[itemIndex].productQuantity += addProductQuantity
				ShowProduct(productList)
			} else {
				productList = append(productList, product)
				ShowProduct(productList)
			}

		case "exit":
			fmt.Println("exit")
		}

	}

}
