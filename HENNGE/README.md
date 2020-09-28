#### Company
    HENNGE
    
#### Position
    Backend Software Engineer
    
### Question
    Write a program which fulfills the requirements below.
    tldr; Parse input without using `for` loop

#### Description
* We want you to calculate the sum of squares of given integers, excluding any negatives.
* The first line of the input will be an integer `N` (`1 <= N <= 100`), indicating the number of test cases to follow.
* Each of the test cases will consist of a line with an integer `X` (`0 < X <= 100`), followed by another line consisting of `X` number of space-separated integers `Yn` (`-100 <= Yn <= 100`).
* For each test case, calculate the sum of squares of the integers, excluding any negatives, and print the calculated sum in the output.
* __Note__: There should be no output until all the input has been received.
* __Note 2__: Do not put blank lines between test cases solutions.
* __Note 3__: Take input from standard input, and output to standard output.

#### Rules
* Write your solution using Go Programming Language
* Your source code must be a single file (`package main`)
* Do not use any `for` statement
* You may only use standard library packages

__Sample Input__
```
2
4
3 -1 1 14
5
9 6 -53 32 16
```

__Sample Output__
```
206
1397
```

Date: 08/07/2020