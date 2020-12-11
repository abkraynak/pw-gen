package main

import(
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("Random Password Generator\n")

	fmt.Printf("Enter the minimum length : ")
	var minLen int
	fmt.Scanf("%d", &minLen)

	fmt.Printf("Enter the maximum length : ")
	var maxLen int
	fmt.Scanf("%d", &maxLen)

	fmt.Printf("Your random password is %v\n", genPassword(minLen, maxLen))
}

func genPassword(minLen int, maxLen int) string {
	// chars used in generated passwords
	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%&*()-=+~?")

	// choose a random length within the parameters
	var length = rand.Intn(maxLen - minLen) + minLen

	fmt.Printf("chars:%v \nlength:%v", chars, length)
	return "x"
}
