package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`
		separator = ","
	)
	rows := strings.Split(data, "\n")
	parsedData := make([][]string, len(rows))
	for i, row := range rows {
		parsedData[i] = strings.Split(row, ",")
	}

	args := os.Args[1:]
	for _, row := range parsedData {
		if len(args) == 0 {
			printColumns(parsedData, row, "")
		} else {
			for _, arg := range args {
				printColumns(parsedData, row, arg)
			}
		}
		fmt.Println()
	}
}

func printColumns(parsedData [][]string, row []string, arg string) {
	for i, column := range row {
		if arg == "" || arg == parsedData[0][i] {
			fmt.Printf("%v\t", column)
		}
	}
}
