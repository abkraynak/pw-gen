/Users/andrewkraynak/go/src/pw-gen package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Random Password Generator\n")

	var len int

	fmt.Printf("Enter desired password length : ")
	fmt.Scanf("%d", &len)

	fmt.Print("Your random password is :")
	genPassword(len)
}

func randomInRange(min int, max int) (int) {
	return rand.Intn(max - min) + min
}

func genPassword(len int) (int) {
	// range of ASCII characters able to type
	minChar := 33
	maxChar := 126

	// seed value to initialize the random number
	SEED := time.Now().Unix()
	rand.Seed(SEED)

	for i := 1; i < len; i++ {
		randChar := randomInRange(minChar, maxChar)
		newChar := string(byte(randChar))
		fmt.Print(newChar)
	}
	fmt.Println()
	return 0
}
