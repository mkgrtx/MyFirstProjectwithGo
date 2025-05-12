package main

import "fmt"

func main() {
	var (
		bubblegum     int = 202
		toffee        int = 118
		iceCream      int = 2250
		milkChocolate int = 1680
		doughnut      int = 1075
		pancake       int = 80
	)
	income := bubblegum + toffee + iceCream + milkChocolate + doughnut + pancake

	// Write your code here
	fmt.Printf(`Earned amount:
Bubblegum: $%d
Toffee: $%d
Ice cream: $%d
Milk chocolate: $%d
Doughnut: $%d
Pancake: $%d

Income: $%d`, bubblegum, toffee, iceCream, milkChocolate, doughnut, pancake, income)
	fmt.Println()
	var staffExp, otherExp int
	fmt.Print("Staff expenses: ")
	fmt.Scanln(&staffExp)
	fmt.Print("Other expenses: ")
	fmt.Scanln(&otherExp)
	netIncome := income - (staffExp + otherExp)
	fmt.Printf("Net income: $%d\n", netIncome)
}
