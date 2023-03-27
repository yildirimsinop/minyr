package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/yildirimsinop/funtemps/conv"
)

func main() {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Filen kunne ikke åpnes:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("CSV kunne ikke leses:", err)
		return
	}

	newFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Kunne ikke opprette ny fil:", err)
		return
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	defer writer.Flush()

	for i, record := range records {
		if i == 0 {
			writer.Write(record)
		} else {
			celsius, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				fmt.Printf("Linje %d ugyldig numerisk verdi: %+v\n", i+1, record)
				continue
			}
			fahrenheit := conv.CelsiusToFahrenheit(celsius)
			record[1] = fmt.Sprintf("%.2f", fahrenheit)
			writer.Write(record)
		}
	}

	conv.CelsiusToFahrenheit(32.0)

	footer := []string{"Data er basert på gyldig data (per 18.03.2023)(CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Umit Yildirim"}
	err = writer.Write(footer)
	if err != nil {
		fmt.Println("Kunne ikke skrive endelig tekst:", err)
	}

}
