package main

import (
	"fmt"
)

type celsius float64
type arrayCelsius []celsius
type fahrenheit float64
type arrayFahrenheit []fahrenheit

func (value fahrenheit) celsius() celsius {
	return celsius((value - 32.0) / 1.8)
}
func (value celsius) fahrenheit() fahrenheit {
	return fahrenheit((value * 1.8) + 32)
}

func (data arrayFahrenheit) dataFtoC() []celsius {
	var newDataCelsius arrayCelsius
	for _, value := range data {
		newDataCelsius = append(newDataCelsius, value.celsius())
	}

	return newDataCelsius
}

func (data arrayCelsius) dataCtoF() []fahrenheit {
	var newDataFahrenheit arrayFahrenheit
	for _, value := range data {
		newDataFahrenheit = append(newDataFahrenheit, value.fahrenheit())
	}

	return newDataFahrenheit
}

func generateDataCelsius(startValue celsius, endValuse celsius, step celsius) []celsius {
	var index_max int = int((endValuse - startValue) / step) + 1
	var newDataCelsius = make([]celsius, 0, index_max)
	var newValueCelsius celsius = 0.0
	for index := 0; index < index_max; index++ {
		newValueCelsius = startValue + (step * celsius(index))
		newDataCelsius = append(newDataCelsius, newValueCelsius)
	}

	return newDataCelsius
}

func generateDataFahrenheit(startValue fahrenheit, endValuse fahrenheit, step fahrenheit) []fahrenheit {
	var index_max int = int((endValuse - startValue) / step) + 1
	var newDataFahrenheit = make([]fahrenheit, 0, index_max)
	var newValueFahrenheit fahrenheit = 0.0
	for index := 0; index < index_max; index++ {
		newValueFahrenheit = startValue + (step * fahrenheit(index))
		newDataFahrenheit = append(newDataFahrenheit, newValueFahrenheit)
	}

	return newDataFahrenheit
}

func drawTable(dataCelsius arrayCelsius, dataFahrenheit arrayFahrenheit) {
	if len(dataCelsius) != len(dataFahrenheit) {
		// Посмотреть как вызывать ошибку
	} else {
		fmt.Println("=======================")
		fmt.Printf("|%10v|%10v|\n", "°C", "°F")
		fmt.Println("=======================")
		for i := 0; i < len(dataCelsius); i++ {
			fmt.Printf("|%10.2f|%10.2f|\n", dataCelsius[i], dataFahrenheit[i])
		}
		fmt.Println("=======================\n")
	}
}

func main() {
	var firstDataC arrayCelsius = generateDataCelsius(40.0, 100.0, 5.0)
	var firstDataF arrayFahrenheit = firstDataC.dataCtoF()

	var secondDataF arrayFahrenheit = generateDataFahrenheit(40.0, 100.0, 5.0)
	var secondDataC arrayCelsius = secondDataF.dataFtoC()

	drawTable(firstDataC, firstDataF)
	drawTable(secondDataC, secondDataF)
}