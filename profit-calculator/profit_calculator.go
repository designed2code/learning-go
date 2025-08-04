package main

import "fmt"

func main(){
	// var revenue float64 
	// var expenses float64
	// var taxRate float64
	revenue := getInputValues("Revenue: ")
	// fmt.Print("Revenue: ") 
	// fmt.Scan(&revenue)

	expenses := getInputValues("Expenses: ")
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	taxRate := getInputValues("TaxRate: ")
	// fmt.Print("TaxRate: ")
	// fmt.Scan(&taxRate)

	
	ebt,profit,ratio := calculateValues(revenue,expenses,taxRate)
	
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func getInputValues(text string) (float64){
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateValues(revenue float64,expenses float64,taxRate float64) (float64,float64,float64){
	ebt := revenue - expenses 
	profit := ebt * (1-taxRate/100)
	ratio := ebt/profit
	return ebt,profit,ratio
}