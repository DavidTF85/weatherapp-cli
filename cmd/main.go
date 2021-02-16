package main

import (
	//go packages
	"fmt"
	"time"

	//"time"


	//my pacakeges:
	"weatherapp-cli/internal/utils"
  "weatherapp-cli/internal/controllers"

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
			fmt.Println("You chosed (A)\n")
			time.Sleep(1 * time.Second)
			controllers.RunAddRecord()

		case "b", "B":
			fmt.Println("(B) List all records")
		case "c", "C":
			fmt.Println("(C) Calculate Summary")
		case "d", "D":
			fmt.Println("(D) Exit")
			return

		default:
			fmt.Println("Hey !!!! FU*&% idot I had sayid choose A,B,C or D.\n")
			fmt.Printf("so read the fu*&@ng intructions and do it again dumass!\n")
		}
	}
}
