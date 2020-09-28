/*
Question:
Please implement the following mathematical function in C++, in the way you
think is best (explain your choice.)

F(0) = 0
F(n) = n + F(n-1)
n=1..20

SOLUTION:

You didn't include other numbers so I am assuming this is NOT the Fibonacci Sequence.
See code below.

*/

#include <iostream>

int F(int n);

int main(int argc, char * argv[]) {
  for (int i = 0; i < 20; i++) {
    printf("F(%i) = %i\n", i, F(i));
  }
}

int F(int n) {
  // Make sure that function doesn't loop forever
  if (n <= 0) {
    return 0;
  }

  return n + F(n - 1);
}
