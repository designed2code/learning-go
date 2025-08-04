package main

import (
	"fmt"
	"math"
)
func main(){
	const inflationRate = 2.5
	var investmentAmount float64
	  expectedReturnRate := 5.5
	 var years float64
	 outputText("Investment Value: ")
	 fmt.Scan(&investmentAmount)
	 outputText("Expected Return Rate: ")
	 fmt.Scan(&expectedReturnRate)
	 outputText("Years: ")
	 fmt.Scan(&years)
	 futureValue := (investmentAmount) * math.Pow(1+expectedReturnRate/100,(years))
	 futureRealValue := futureValue/math.Pow(1+inflationRate/100,years)

	 formattedFV := fmt.Sprintf("Future Value %.2f\n",futureValue)
	// fmt.Println("Future Value",futureValue)
	fmt.Print(formattedFV)
	fmt.Println("Future Value Adjusted for inflation",futureRealValue)
}

func outputText(text string){
	fmt.Print(text)
}