package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
	    reader := bufio.NewReader(os.Stdin)
	    numbers, _ := reader.ReadString('\n')

	    tokenized := strings.Split(numbers, ",")
	    var newList []string
	    removed := make(map[int]bool)

	    for _, number := range tokenized {
		        i, err := strconv.Atoi(number)
		        if err != nil {
			            fmt.Printf(err.Error())
			            continue
			        }

		        if i % 2 == 0 || removed[i] == true {
            newList = append(newList, number)
        } else {
            removed[i] = true
        }
    }

    fmt.Println(strings.Join(newList, ","))
}