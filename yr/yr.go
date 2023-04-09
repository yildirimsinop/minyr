package yr

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	conv "github.com/yildirimsinop/funtemps/Conv"
)

func ConvertTemperature() {
	overwriteFile := checkFileExists()
	if !overwriteFile {
		fmt.Println("Går tilbake til hovedmeny")
		return
	}

	inputFile := openInputFile()
	defer inputFile.Close()

	outputFile, err := createOutputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	scanner := bufio.NewScanner(inputFile)

	if scanner.Scan() {
		_, err := outputWriter.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		outputLine := ProcessLine(line)
		_, err := outputWriter.WriteString(outputLine + "\n")
		if err != nil {
			panic(err)
		}
	}

	outputWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ferdig!")
}

func checkFileExists() bool {
	if _, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
		fmt.Print("Filen eksisterer allerede. Vil du generere filen på nytt? (y/n): ")
		var overwriteInput string
		fmt.Scanln(&overwriteInput)
		if strings.ToLower(overwriteInput) == "y" {
			err := os.Remove("kjevik-temp-fahr-20220318-20230318.csv")
			if err != nil {
				log.Fatal(err)
			}
			return true
		}
		return false
	}
	return true
}

func openInputFile() *os.File {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func createOutputFile() (*os.File, error) {
	outputFilePath := "kjevik-temp-fahr-20220318-20230318.csv"
	if _, err := os.Stat(outputFilePath); err == nil {
		fmt.Printf("File %s already exists. Deleting...\n", outputFilePath)
		err := os.Remove(outputFilePath)
		if err != nil {
			return nil, fmt.Errorf("could not delete file: %v", err)
		}
	}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not create file: %v", err)
	}
	return outputFile, nil
}

func ProcessLine(line string) string {
	if line == "" {
		return ""
	}
	fields := strings.Split(line, ";")
	lastField := ""
	if len(fields) > 0 {
		lastField = fields[len(fields)-1]
	}
	convertedField := ""
	if lastField != "" {
		var err error
		convertedField, err = convertLastField(lastField)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return ""
		}
	}
	if convertedField != "" {
		fields[len(fields)-1] = convertedField
	}
	if line[0:7] == "Data er" {
		return "Data er basert på gyldig data (per 18.03.2023)(CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Umit Yildirim"
	} else {
		return strings.Join(fields, ";")
	}
}

func convertLastField(lastField string) (string, error) {
	celsius, err := strconv.ParseFloat(lastField, 64)
	if err != nil {
		return "", err
	}

	fahrenheit := conv.CelsiusToFahrenheit(celsius)

	return fmt.Sprintf("%.1f", fahrenheit), nil
}

func AverageTemperature() {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Velg temperaturenhet (celsius/fahr):")
	var unit string
	fmt.Scan(&unit)

	var sum float64
	count := 0
	for i, line := range lines {
		if i == 0 {
			continue
		}
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			log.Fatalf("unexpected number of fields in line %d: %d", i, len(fields))
		}
		if fields[3] == "" {
			continue
		}
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			log.Fatalf("could not parse temperature in line %d: %s", i, err)
		}

		if unit == "fahr" {
			temperature = conv.CelsiusToFahrenheit(temperature)
		}
		sum += temperature
		count++
	}

	if unit == "fahr" {
		average := sum / float64(count)
		average = math.Round(average*100) / 100
		fmt.Printf("Gjennomsnittlig temperatur: %.2f°F\n", average)
	} else {
		average := sum / float64(count)
		fmt.Printf("Gjennomsnittlig temperatur: %.2f°C\n", average)
	}
}
func CountLines(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	countedLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			countedLines++
		}
	}
	return countedLines
}

func GetAverageTemperature(filepath string, unit string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var sum float64
	count := 0
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			continue
		}
		fields := strings.Split(scanner.Text(), ";")
		if len(fields) != 4 {
			return "", fmt.Errorf("unexpected number of fields in line %d: %d", i, len(fields))
		}
		if fields[3] == "" {
			continue
		}
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return "", fmt.Errorf("could not parse temperature in line %d: %s", i, err)
		}

		if unit == "fahr" {
			temperature = conv.CelsiusToFahrenheit(temperature)
		}
		sum += temperature
		count++
	}
	average := sum / float64(count)
	return fmt.Sprintf("%.2f", average), nil
}
