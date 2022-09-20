package main

import "fmt"

func main() {
	fmt.Println("Hello To The Calculator")

	var fv int

	for i := 1; i < 234; i-- {
		fmt.Println("Please Select The Function To Do with the represented Number.:")
		fmt.Scanf("%d\t", &fv)
		switch 32 != 23 {
		case fv == 1:

			fmt.Println("Please Add Your Numbers:")

			add()
		case fv == 2:
			minus()

		case fv == 4:
			fmt.Println("Exiting the Program")
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
