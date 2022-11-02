package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello To The Calculator")

	var fv int
	fmt.Println("Hi")

	for i := 1; i < 234; i-- {
		fmt.Println("Please Select The Function To Do with the represented Number.:")
		fmt.Println("1. Addition  2. Minus  3. Multiplication  4. Divison   5. Exit the program.")
		fmt.Scanf("%d\t", &fv)
		switch {
		case fv == 1:

			fmt.Println("Please Add Your Numbers:")

			add()
		case fv == 2:
			minus()

		case fv == 3:
			multi()

		case fv == 4:
			Divison()
		case fv == 5:
			fmt.Println("Exited the Program")

			return
		default:
			fmt.Println("Please Enter a Valid number you retard.")

		}

	}
}

func add(x ...int) {
	var num1, num2 int
	fmt.Print("Enter the First Number = ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the Second Number = ")
	fmt.Scanln(&num2)

	fmt.Println("The Sum of num1 and num2  = ", num1+num2)

}
func minus() {
	var num1, num2 int
	fmt.Print("Enter the First Number = ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the Second Number = ")
	fmt.Scanln(&num2)

	fmt.Println("The Minus of num1 and num2  = ", num1-num2)

}
func multi() {
	var num1, num2 float64
	fmt.Println("Please Add in your numbers:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Println(num1, "into", num2, "is :", num1*num2)

}
func Divison() {
	var num1, num2 float64
	fmt.Println("Please Add in your numbers:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Println(num1, "Divided by", num2, "= ", num1/num2)
}
