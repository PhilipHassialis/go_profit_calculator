package main

import "fmt"

func main() {

	var revenue, expenses, taxRate float64

	userInput("Enter revenue:", &revenue)
	userInput("Enter expenses:", &expenses)
	userInput("Enter tax rate:", &taxRate)

	earningsBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

	printOutput(earningsBeforeTax, profit, ratio)
}

func printOutput(earningsBeforeTax, profit, ratio float64) {
	fmt.Printf("EBT: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func userInput(userPrompt string, userValue *float64) {
	fmt.Print(userPrompt, " ")
	fmt.Scan(userValue)
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit
	return earningsBeforeTax, profit, ratio
}
