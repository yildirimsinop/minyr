package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convert() {
	if len(os.Args) != 2 || os.Args[1] != "minyr" {
		fmt.Println("Usage: minyr")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter 'convert' or 'average'")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "convert":
		// check if file exists
		if _, err := os.Stat("kjevik-tempfahr-20220318-20230318.csv"); os.IsNotExist(err) {
			// file does not exist, generate it
			fmt.Println("Generating file...")
			// generateFile()
		} else {
			// file exists, ask user if they want to regenerate it
			fmt.Println("File already exists. Do you want to regenerate it? (y/n)")
			regenerate, _ := reader.ReadString('\n')
			regenerate = strings.TrimSpace(regenerate)
			if regenerate == "y" {
				// regenerateFile()
				fmt.Println("File regenerated.")
			} else {
				fmt.Println("File not regenerated.")
			}
		}
	case "average":
		// read file and calculate average temperature
		// averageTemp := calculateAverageTemp("kjevik-tempfahr-20220318-20230318.csv")
		// fmt.Println("Average temperature:", averageTemp)

		// ask user if they want Celsius or Fahrenheit
		fmt.Println("Do you want Celsius or Fahrenheit? (c/f)")
		tempChoice, _ := reader.ReadString('\n')
		tempChoice = strings.TrimSpace(tempChoice)

		if tempChoice == "c" {
			// calculate and print temperature in Celsius
			// tempCelsius := convertToFahrenheit(averageTemp)
			// fmt.Println("Average temperature in Celsius:", tempCelsius, "°C")
		} else if tempChoice == "f" {
			// calculate and print temperature in Fahrenheit
			// fmt.Println("Average temperature in Fahrenheit:", averageTemp, "°F")
		} else {
			fmt.Println("Invalid choice.")
		}
	default:
		fmt.Println("Invalid choice.")
	}
}

// func generateFile() {
//     // generate file
// }

// func calculateAverageTemp(filename string) float64 {
//     // read file and calculate average temperature
//     return 0.0
// }

// func convertToCelsius(tempF float64) float64 {
//     // convert temperature from Fahrenheit to Celsius
//     return 0.0
// }

// func convertToFahrenheit(tempC float64) float64 {
//     // convert temperature from Celsius to Fahrenheit
//     return 0.0
// }
