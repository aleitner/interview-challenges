/*
Question:
Please think of an algorithm that given 4 arbitrary points produces a smooth
curve that connects all of them.Then write a function in C++ that takes 5
parameters ( 4 points and a "time" parameter between 0 and 1) and returns
the point on the curve at an arbitrary "time".

(9b)

Now write a C++ function that given an arbitrary number of points smoothly
interpolates between them, making use of the function/functions you wrote in
9a if appropriate.
Some form of visual output is a plus.

SOLUTION:
Lagrange algorithm?
*/

#include <iostream>

struct Point {
  int x, y;
};

// 9b
double interpolate(struct Point * points, int xi, int n) {
  double result = 0;

  for (int i = 0; i < n; i++) {
    double term = points[i].y;

    for (int j = 0; j < n; j++) {
      if (j != i) {
        term = term * (xi - points[j].x) / double(points[i].x - points[j].x);
      }
    }

    result += term;
  }

  return result;
}

// 9a
double interpolate_range(struct Point * points, double xi, int known_points_total) {
  int min = points[0].x;
  int max = points[known_points_total - 1].x;
  printf("min = %d, max = %d, known_points_total: %d\n", min, max, known_points_total);

  int scaled_xi = min + (double(max - min) * xi);

  printf("x = %d\n", scaled_xi);

  return interpolate(points, scaled_xi, known_points_total);
}

int main() {
  Point points[] = {{0,2}, {1,3}, {2,12}, {5,147}};

  size_t known_points_total = (int) sizeof(points) / sizeof(Point);

  double scale = 0.2;
  printf("Value of f(%d%%) is : %f", int(scale * 100), interpolate_range(points, scale, known_points_total));
}
