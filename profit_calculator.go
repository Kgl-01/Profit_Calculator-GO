package main

import (
	"errors"
	"fmt"
)

func main() {

	revenue, err1 := readUserInput("Revenue: ")

	expenses, err2 := readUserInput("Expense: ")

	taxRate, err3 := readUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateFinancialsAlt(revenue, expenses, taxRate)

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

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

/*Alternative return snyntax for function*/
func calculateFinancialsAlt(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	//return ebt, profit, ratio (or)
	return
}

func readUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number")
	}
	return userInput, nil
}
