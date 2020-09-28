package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MathMachine performs calculations on a number of test cases
type MathMachine struct {
	step         int         // Current step in the Machine. We step every time we read a testcase from stdin
	numTestCases int         // Number of test cases to be evaluated
	testCases    []*TestCase // slice of structs containing test cases and calculations
	printStep    int         // Current step for printing the results
}

// PrintResults prints the results of the TestCase calculations
func (mm *MathMachine) PrintResults() error {
	if mm.printStep == mm.numTestCases {
		mm.printStep = 0
		return nil
	}

	fmt.Println(mm.testCases[mm.printStep].result)

	mm.printStep++

	return mm.PrintResults()
}

// IsFinished checks if the mathmachine has no more test cases
func (mm MathMachine) IsFinished() bool {
	// Check if we can end only after we have received our number of test cases
	if mm.step > 0 {
		// If we have no tests then end
		if mm.numTestCases == 0 {
			return true
		}

		// If we have stepped through all the test cases then end
		if mm.step > mm.numTestCases {
			return true
		}
	}

	return false
}

// setNumTestCases sets the number of testcases to be calculated
func (mm *MathMachine) setNumTestCases() error {
	// Only the first step can set the number of test cases
	if mm.step != 0 {
		return nil
	}

	i, err := readNextLineAsInt()
	if err != nil {
		return err
	}

	mm.numTestCases = i
	mm.step++
	return nil
}

// getTestCase reads the next testcase related lines from stdin and package them into a TestCase struct
func (mm *MathMachine) getTestCase() (*TestCase, error) {
	numInts, err := readNextLineAsInt()
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(os.Stdin)
	intList, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	mm.step++

	return &TestCase{
		numInts: numInts,
		intList: strings.Split(intList, " "),
	}, nil
}

// Prepare reads all the testcases from stdin and runs calculations on them
func (mm *MathMachine) Prepare() error {
	// Check if we are finished
	if mm.IsFinished() {
		return nil
	}

	// step 0 determines the number of test cases
	if err := mm.setNumTestCases(); err != nil {
		return err
	}

	// get the next testcase
	tc, err := mm.getTestCase()
	if err != nil {
		return err
	}

	mm.testCases = append(mm.testCases, tc)

	// perform calculations on the testcase
	if err := tc.calculate(); err != nil {
		return err
	}

	return mm.Prepare()
}

// TestCase contains the data to be calculated
type TestCase struct {
	step    int
	numInts int
	intList []string
	result  int
}

// calculate the result for the TestCase
func (tc *TestCase) calculate() error {
	// check if finished calculations
	if tc.step == tc.numInts {
		return nil
	}

	// Get current number to be calculated from the list based on the step
	currentNumAsString := strings.TrimSuffix(tc.intList[tc.step], "\n")
	currentNum, err := strconv.Atoi(currentNumAsString)
	if err != nil {
		return err
	}

	if currentNum > 0 {
		tc.result += currentNum * currentNum
	}

	tc.step++

	return tc.calculate()
}

func main() {
	tracker := &MathMachine{}
	if err := tracker.Prepare(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := tracker.PrintResults(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// readNextLineAsInt will read a line from stdin and convert it to an integer
func readNextLineAsInt() (num int, err error) {
	var i int
	_, err = fmt.Scanf("%d", &i)

	return i, err
}
