package utils

import (
	"fmt"
	"log"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func CreateBill(name string) Bill {
	newBill := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return newBill
}

func (bill Bill) format() string {

	formattedText := fmt.Sprintf("%v", "- - - - - - - THE JASON HOUSE - - - - - - -\n")

	var total float64 = 0

	for key, value := range bill.items {
		formattedText += fmt.Sprintf("%-25v ...$%v\n", key+":", value)
		total += value
	}
	total += bill.tip

	formattedText += fmt.Sprintf("%-25v ...$%v\n", "Tip:", bill.tip)
	formattedText += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total)

	return formattedText
}

func (bill *Bill) UpdateTip(tip float64) {
	bill.tip = tip
}

func (bill Bill) AddItem(name string, price float64) {
	bill.items[name] = price
}

func exportBill(bill Bill) {
	data := []byte(bill.format())

	path := "bills/" + bill.name + ".txt"

	error := os.WriteFile(path, data, 0644)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Your bill was saved to a file!")
}

func (bill Bill) SaveBill() {
	path := "bills/" + bill.name + ".txt"

	// Create a folder when it doesn't exist.
	if _, err := os.Stat("bills/"); os.IsNotExist(err) {
		os.MkdirAll("bills/", 0700)
	}

	// Replace a file when it exists.
	if _, err := os.Stat(path); err == nil {
		error := os.Remove(path)

		if error != nil {
			log.Fatal(error)
		}

		exportBill(bill)

	} else {
		exportBill(bill)
	}
}
