package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	formattedEbt := fmt.Sprintf("Expense before tax: %.2f \n", ebt)
	formattedProfit := fmt.Sprintf("Your profit is: %.2f \n", profit)
	formattedRatio := fmt.Sprintf("EBT / Profit: %.2f \n", ratio)

	fmt.Print(formattedEbt)
	fmt.Print(formattedProfit)
	fmt.Print(formattedRatio)
}
