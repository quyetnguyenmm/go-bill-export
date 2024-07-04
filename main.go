package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(reader *bufio.Reader, promt string) (string, error) {
	fmt.Print(promt)
	name, error := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return name, error
}

func promtOptions(bill Bill) {
	fmt.Println("\n- - - - - - - Options - - - - - - -")
	fmt.Println("a - Add items")
	fmt.Println("t - Add tip")
	fmt.Println("s - Save bill")
	fmt.Print("")

	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput(reader, "Please choose your option: ")

	switch option {
	case "a":
		name, _ := getInput(reader, "Item name: ")
		price, _ := getInput(reader, "Item price: ")

		parsedPrice, error := strconv.ParseFloat(price, 64)
		if error != nil {
			fmt.Println("The price must be a number!")
			promtOptions(bill)
		}

		bill.addItem(name, parsedPrice)
		fmt.Println("The item was added successfully!")
		promtOptions(bill)
	case "t":
		tip, _ := getInput(reader, "Tip amount ($): ")

		parsedTip, error := strconv.ParseFloat(tip, 64)
		if error != nil {
			fmt.Println("The tip must be a number!")
			promtOptions(bill)
		}

		bill.updateTip(parsedTip)
		fmt.Println("The tip was updated successfully!")
		promtOptions(bill)
	case "s":
		bill.saveBill()
	default:
		fmt.Println("Your option is not valid! Choose again")
		promtOptions(bill)
	}
}

func createBillFromUser() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput(reader, "Please enter your bill name: ")

	newBill := createBill(name)
	fmt.Println("The bill was created - ", name)

	return newBill
}

func main() {
	bill := createBillFromUser()
	promtOptions(bill)
}
