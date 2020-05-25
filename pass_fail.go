package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
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
