package main

import "fmt"

func main(){
	var revenue float64 
	var expenses float64
	var taxRate float64
	fmt.Print("Revenue: ") 
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("TaxRate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses 
	profit := ebt * (1-taxRate/100)
	ratio := ebt/profit
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}