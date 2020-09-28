/*
Question:
Please write a program that displays times tables in C++.
*/

#include <iostream>

int main() {
  int max = 12;

  for (int y = 1; y <= max; y++) {
    printf("|");

    for (int x = 1; x <= max; x++) {
      printf("%4d", x * y);
      printf("|");
    }

    printf("\n");
  }
}
