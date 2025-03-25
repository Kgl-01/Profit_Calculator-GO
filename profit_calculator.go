package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("Revenue: ")
	fmt.Scan(&revenue)

	outputText("Expenses: ")
	fmt.Scan(&expenses)

	outputText("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateEbtProfitAndRatioAlt(revenue, expenses, taxRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	formattedEbt := fmt.Sprintf("Expense before tax: %.2f \n", ebt)
	formattedProfit := fmt.Sprintf("Your profit is: %.2f \n", profit)
	formattedRatio := fmt.Sprintf("EBT / Profit: %.2f \n", ratio)

	//Single line string formatted values
	fmt.Print(formattedEbt)
	fmt.Print(formattedProfit)
	fmt.Print(formattedRatio)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateEbtProfitAndRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

/*Alternative return snyntax for function*/
func calculateEbtProfitAndRatioAlt(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	//return ebt, profit, ratio (or)
	return
}
