package main

import (
	"fmt"
	"time"
)

func fromEden() {

	// Initializing time (sec)
	one_sec := 1 * time.Second

	line1 := "Honey, "
	line2 := "you're familiar "
	line3 := "like my mirror years ago "

	line4 := "Idealism "
	line5 := "sits in prison "
	line6 := "chivalry fell on its sword "

	line7 := "Innonence "
	line8 := "died screaming, "
	line9 := "honey, ask me I should know "

	line10 := "I slithered here "
	line11 := "from Eden "
	line12 := "just to hide outside your door "

	fmt.Println()
	fmt.Println()

	_type(line1)
	time.Sleep(one_sec)
	_type(line2)
	time.Sleep(one_sec + 40*time.Millisecond)
	_type(line3)

	fmt.Println()
	time.Sleep(one_sec + one_sec)

	_type(line4)
	time.Sleep(one_sec)
	_type(line5)
	time.Sleep(one_sec + 40*time.Millisecond)
	_type(line6)

	fmt.Println()
	time.Sleep(one_sec + one_sec)

	_type(line7)
	time.Sleep(one_sec)
	_type(line8)
	time.Sleep(one_sec + 40*time.Millisecond)
	_type(line9)

	fmt.Println()
	time.Sleep(one_sec + one_sec)

	_type(line10)
	time.Sleep(one_sec)
	_type(line11)
	time.Sleep(one_sec + 40*time.Millisecond)
	_type(line12)

	fmt.Println()
	fmt.Println()
	fmt.Println(":)")

}

func _type(line string) {
	delay := 50 * time.Millisecond
	for _, char := range line {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
}

func main() {

	fromEden()

}
