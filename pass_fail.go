package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatalln(err)
	}

	status := "failing"
	if grade >= 60 {
		status = "passing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
