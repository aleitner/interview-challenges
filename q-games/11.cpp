/*
Question:
Think of a question that you feel should be on this test and answer it.
(please don't skip that question)

Implement quesiton 10 for C++

SOLUTION: See Below

Things that can be optimized:
- Group things into quadrants and only check local quadrants
- Optimize distance method for checking different shaped Things
*/

#include <cmath>
#include <cstdlib>
#include <iostream>
#include <time.h>

int distance(int x1, int y1, int x2, int y2);

class Thing {
  public:
    int x, y, r;
  Thing() {};
  Thing(int, int, int);
  int area() {
    return M_PI * pow(r, 2);
  };
  bool collides(Thing);
};

Thing::Thing(int xpos, int ypos, int radius) {
  x = xpos;
  y = ypos;
  r = radius;
}

bool Thing::collides(Thing thing2) {
  int d = distance(x, y, thing2.x, thing2.y);

  if (d < r + thing2.r) {
    return true;
  }

  return false;
}

// Calculate distance between two points
// TODO: We need change this depending on the shape of the Thing
int distance(int x1, int y1, int x2, int y2) {
  return sqrt(pow((x1 - x2), 2) + pow((y1 - y2), 2));
}

int main() {
  srand(time(NULL));

  int thing_count = 5000;

  Thing * things;
  things = new Thing[thing_count];

  // Create everything on the screen
  for (int i = 0; i < thing_count; i++) {
    things[i] = Thing(rand() % 5000, rand() % 5000, 1);
  }

  // check collisions by looping through everything on screen
  for (int i = 0; i < thing_count; i++) {
    for (int j = i; j < thing_count; j++) {
      if (i != j && things[i].collides(things[j]) == true) {
        printf("%d collides with %d\n", i, j);
      }
    }
  }

}
