package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yildirimsinop/minyr/yr"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Venligst velg convert, average eller exit:")

		if !scanner.Scan() {
			break
		}
		input = scanner.Text()

		switch input {
		case "q", "exit":
			fmt.Println("exit")
			return

		case "convert":
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit...")
			yr.ConvertTemperature()

		case "average":
			fmt.Println("Gjennomsnitt-kalkulator")

			for {

				yr.AverageTemperature()

				var input2 string
				scanjn := bufio.NewScanner(os.Stdin)
				fmt.Println("Tilbake til hovedmeny? (y/n)")
				for scanjn.Scan() {
					input2 = scanjn.Text()
					if input2 == "y" {
						break
					} else if input2 == "n" {
						break
					}
				}
				if input2 == "y" {
					break
				}
			}
		}
	}
	fmt.Println("Avslutter program.")
}
