package main

import (
	"fmt"
	"math"
)
const inflationRate = 2.5
func main(){
	var investmentAmount float64
	  expectedReturnRate := 5.5
	 var years float64
	 outputText("Investment Value: ")
	 fmt.Scan(&investmentAmount)
	 outputText("Expected Return Rate: ")
	 fmt.Scan(&expectedReturnRate)
	 outputText("Years: ")
	 fmt.Scan(&years)
	 futureValue,futureRealValue := calculateFutureValues(investmentAmount,expectedReturnRate,years)
	 

	 formattedFV := fmt.Sprintf("Future Value %.2f\n",futureValue)
	// fmt.Println("Future Value",futureValue)
	fmt.Print(formattedFV)
	fmt.Println("Future Value Adjusted for inflation",futureRealValue)
}

func outputText(text string){
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64,expectedReturnRate float64,years float64) (float64,float64){

	fv := (investmentAmount) * math.Pow(1+expectedReturnRate/100,(years))
	rfv := fv/math.Pow(1+inflationRate/100,years)

	return fv,rfv
}