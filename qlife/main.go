package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read from stdin
	// Comment this out because we use custom input data
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')

	// Use custom input since we are
	input := `subject1,subject2,subject3,subject4
98,99,97,100
86,98,56,75
93,55,98,68
45,78,78,89`

	// Read CSV data from input string
	r := csv.NewReader(strings.NewReader(input))

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	headers := rows[0]
	numDataRows := len(rows) - 1
	numColumns := len(headers)

	averages := make([]int, numColumns)
	averagesStr := make([]string, numColumns)

	// Sum each column
	for c := 0; c < numColumns; c++ {
		for r := 1; r <= numDataRows; r++ {

			grade := rows[r][c]
			gradeNum, err := strconv.Atoi(grade)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			averages[c] += gradeNum
		}
	}

	// Average the data and convert to string
	for i, _ := range averages {
		averages[i] /= 4
		averagesStr[i] = strconv.Itoa(averages[i])
	}

	// Write csv data
	csvwriter := csv.NewWriter(os.Stdout)
	if err := csvwriter.Write(headers); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := csvwriter.Write(averagesStr); err != nil {
		fmt.Println(err.Error())
		return
	}
	csvwriter.Flush()
}
