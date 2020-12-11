package main

import(
	"fmt"
)

func main() {
	fmt.Printf("Random Password Generator\n")

	fmt.Printf("Enter the minimum length : ")
	var minLen int
	fmt.Scanf("%d", &minLen)

	fmt.Printf("Enter the maximum length : ")
	var maxLen int
	fmt.Scanf("%d", &maxLen)

	fmt.Printf("You entered minLen=%v and maxLen=%v \n", minLen, maxLen)
}
