package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	// messageVal := *message
	// fmt.Println("--------------")
	// fmt.Println(messageVal)
	// fmt.Println()
	// fmt.Println(*message)
	// fmt.Println()
	// messageVal = messageVal + "modified"
	// fmt.Println(messageVal)
	// fmt.Println()
	// fmt.Println("=============")
	*message = strings.ReplaceAll(*message, "dang", "****")
	*message = strings.ReplaceAll(*message, "shoot", "*****")
	*message = strings.ReplaceAll(*message, "heck", "****")
	// fmt.Println(*message)
}

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}
