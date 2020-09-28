package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "unicode"
)

func main() {
	    reader := bufio.NewReader(os.Stdin)
	    sentence, _ := reader.ReadString('\n')

	    _, resultNum := Count(sentence)

	    fmt.Println(resultNum)
	    //Enter your code here. Read input from STDIN. Print output to STDOUT
}

// Count returns capitalized words and numbers from a string
func Count(sentence string) (words map[string]bool, numWords int) {
	    words = make(map[string]bool)
	    tokenized := strings.Fields(sentence)

	    for _, k := range tokenized {
		        // TODO: strip symbols from k
		        if (isCapitalized(k) || isNumber(k)) {
            words[k] = true
        }
    }

    return words, len(words)
}

// isCapitalized returns if the first letter of a string is capitalized
func isCapitalized(s string) bool {
	    // We have to loop to get the rune
	    for _, ch := range s {
		        if unicode.IsUpper(ch) {
			            return true
			        } else {
		            return false
		        }
		    }

	    return false
}

// isNumber checks if a string is a number
func isNumber(s string) bool {
	    if _, err := strconv.ParseInt(s,10,64); err != nil {
		        return false
		    }

	    return true
}

