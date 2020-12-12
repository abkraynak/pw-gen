package main

import(
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Random Password Generator\n")

	var len int

	fmt.Printf("Enter desired password length : ")
	fmt.Scanf("%d", &len)

	fmt.Print("Your random password is : ")
	genPassword(len)
}

func randomInRange(min int, max int) (int, error) {
	if max - min <= 0 {
		return -1, errors.New("Range of values is zero or negative")
	}
	return rand.Intn(max - min) + min, nil	
}

func genPassword(len int) (int) {
	// range of ASCII characters able to type
	minChar := 33
	maxChar := 126

	// seed value to initialize the random number
	SEED := time.Now().Unix()
	rand.Seed(SEED)

	for i := 1; i < len; i++ {
		randChar, _ := randomInRange(minChar, maxChar)
		newChar := string(byte(randChar))
		fmt.Print(newChar)
	}
	fmt.Println()
	return 0
}
