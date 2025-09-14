package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measurements struct{
	Name string
	Min float64
	Max float64
	Sum float64
	Count int64
}

func main(){
	measurment, err := os.Open("measurements.txt")
	if err != nil{
		panic(err)
	}
	defer measurment.Close()

	dados := make(map[string]Measurements)

	scanner := bufio.NewScanner(measurment)
	for scanner.Scan(){
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurment, ok := dados[location]
		if !ok {
			measurment = Measurements{
				Min: temp,
				Max: temp,
				Sum: temp,
				Count: 1,
			}
		} else {
			measurment.Name = location
			measurment.Min = min(measurment.Min, temp)
			measurment.Max = max(measurment.Min, temp)
			measurment.Sum += temp
			measurment.Count ++
		}

		dados[location] = measurment
	}


	for name, measurment := range dados {
		fmt.Printf("&s: %#+v\n", name, measurment)
	}
}