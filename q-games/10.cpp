/*
Question:
Assuming that you have a 2D game that can display hundreds of non-rotating
sprites. Explain how you could efficiently optimize collision detection on
these sprites without having to test every sprite against every other
sprite.

SOLUTION: See Below
*/

#include <iostream>

int main() {
  printf("1.\tLoop through all objects on screen\n");

  printf("2.\tCreate a loop through every object on screen inside part 1.\n");
  printf("  \tMake sure the loop starting point is incremented every time \n");
  printf("  \tloop 1 iterates. in order to prevent comparing values that \n");
  printf("  \thave already been compared\n");

  printf("3.\tCheck the distances of the two objects' points.\n");

  printf("\n\tSee question 11 for implementation.\n");
}
