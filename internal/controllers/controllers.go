package controllers


import (
	"fmt"
	"time"
	"weatherapp-cli/internal/models"
	"weatherapp-cli/internal/utils"
)
// The functionality for "Add", "List", etc need to be called
//from the weatherapp-cli/internal/controllers/controller.go.
//THIS IS IMPORTANT.

//func RunAddRecord() models.TimeSeriesDatum {

func RunAddRecord() {
    
	//gui for add recods
	fmt.Println("Ok! lets add some recods")
	time.Sleep(2 * time.Second)
	fmt.Println("\n")
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	fmt.Println("\n")
	fmt.Println("Please enter the following records:\n")


	fmt.Println("Temperature (deg):")
	tem:= utils.ReadConsoleFloat64()
	fmt.Println("\n")

	fmt.Println("Humidity (%):")
	Hu:= utils.ReadConsoleFloat64()
	fmt.Println("\n")


	fmt.Println("Pressure (Pa):")
    Pr:= utils.ReadConsoleFloat64()
	fmt.Println("\n")

	fmt.Println("CO2 (ppm):")
	co2:= utils.ReadConsoleFloat64()
	fmt.Println("\n")

	fmt.Println("TVOC (ppb):")
	tvoc:= utils.ReadConsoleFloat64()
	fmt.Println("\n")

	// fmt.Println("Date/Time:")
	// fmt.Println("\n")

	
//data structure
	
newdata:= models.TimeSeriesDatum {
		Temperature: tem,
		Pressure:	Pr,
		Humidity: Hu,
		CarbonDioxide: co2,
		TotalVolatileOrganicCompounds: tvoc,
	}
	fmt.Println(newdata)
// 	//return newdata
}

// //ShowNewData is for create an array and storage new input data 
// func ShowNewData(neoData []models.TimeSeriesDatum) {
// 	fmt.Println("###############################")
// 	fmt.Println("# Weather App Command Console #")
// 	fmt.Println("###############################")
// 	for _, v := range neoData {
// 		fmt.Println("")
// 		fmt.Printf("new Temp:%f \n", v.Temperature)
// 		fmt.Printf("new Pre:%f \n", v.Pressure)
// 		fmt.Printf("new Humi:%f \n", v.Humidity)
// 		fmt.Printf("new CO2:%f \n", v.CarbonDioxide)
// 		fmt.Printf("new Tvoc :%f \n", v.TotalVolatileOrganicCompounds)
// 		fmt.Println("")
// 	}
// 	fmt.Println("this  is all F)(*&Yuo9-")
// } 
