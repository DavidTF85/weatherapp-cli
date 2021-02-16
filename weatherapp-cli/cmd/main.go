package main

import (
	//go packages
	"fmt"
	"time"

	//"time"

	//my pacakeges:
	"weatherapp-cli/internal/utils"
	///asddasdasdassasaasdasd
)

func main() {

	fmt.Println("\n")
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	fmt.Println("\n")
	fmt.Println("Menu")
	fmt.Println("(A) Add a record") //asasdas
	fmt.Println("(B) List all records")
	fmt.Println("(C) Calculate Summary")
	fmt.Println("(D) Exit\n")
	fmt.Println("Please enter choice:")
	fmt.Println("")

	for {

		choice := utils.ReadConsoleString()
		switch choice {
		case "a", "A":
			fmt.Println("You chosed (A)")
			time.Sleep(2 * time.Second)
			fmt.Println("Ok! lets add some recods")
			time.Sleep(3 * time.Second)
			runAddRecord()

		case "b", "B":
			fmt.Println("(B) List all records")
		case "c", "C":
			fmt.Println("(C) Calculate Summary")
		case "d", "D":
			fmt.Println("(D) Exit")
			return

		default:
			fmt.Println("Hey !!!! Fu*&% idot hay say choose A,B,C or D. so read the fu*&@ng intructions and do it again dumass!\n")
		}
	}
}
