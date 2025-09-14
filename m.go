package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Measurement struct{
	Nome string
	Min float64
	Max float64
	Sum float64
	Count int64
}

func main(){
	measurments, err := os.Open("measurements.txt")
	if err != nil{
		panic(err)
	}

	defer measurments.Close()

	scanner := bufio.NewScanner(measurments)
	for scanner.Scan(){
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		temp := rawData[semicolon+1:]

		fmt.Println(location, temp)
		return
	}
}