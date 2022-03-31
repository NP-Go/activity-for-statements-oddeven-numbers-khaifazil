package main

import "fmt"

func main() {
	//Insert code here
	for {
		var num int
		fmt.Println("Enter number: ")
		fmt.Scanln(&num)

		if num%2 != 0 {
			fmt.Println("Number is Odd!")
		} else {
			fmt.Println("Number is Even!")
		}
	}
}
