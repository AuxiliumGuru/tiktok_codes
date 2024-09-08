package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {

	targetWord := "hello world"
	validChars := "abcdefghijklmnopqrstuvwxyz "

	var result string

	for _, i := range targetWord {

		isFound := false
		for !isFound {
			// Generate a random number between 0 and len(validChars)
			randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(validChars))))
			if err != nil {
				panic(err)
			}

			randIndex := randomNum.Int64()

			randChar := string(validChars[randIndex])

			if string(i) != string(randChar) {
				fmt.Println(result + randChar)
			} else {
				result = result + string(i)
				fmt.Println(result)
				isFound = true
			}

			time.Sleep(20 * time.Millisecond)

		}

	}
}
