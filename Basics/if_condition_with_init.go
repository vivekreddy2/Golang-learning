package main

import "fmt"

// init statement - the scope of text variable is only within if else block


func main() {
	
	if text := "Golang"; len(text) > 10 {
		fmt.Println("Yes,", text)
	} else {
		fmt.Println("No,", text)
	}
}
