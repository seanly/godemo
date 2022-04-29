package hello

import "fmt"

func Hello() {
	fmt.Println("hello func")
}

var Util int = 5

func init() {
	fmt.Println("init test main")

	Hello()
}

