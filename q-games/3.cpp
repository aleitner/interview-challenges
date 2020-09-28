/*
Question:
Please use the STL vector class to make a triangle list class, for use in 3d
drawing. For each triangle please store 3 vertex positions, 3 colours, and 1
face normal.
*/

#include <cmath>
#include <cstdlib>
#include <iostream>
#include <vector>
using namespace std;

class Point {
  public:
    double x, y, z;

    Point() {
      x = 0.0;
      y = 0.0;
      z = 0.0;
    };

    Point(double, double, double);
};

Point::Point(double xpos, double ypos, double zpos) {
  x = xpos;
  y = ypos;
  z = zpos;
};

class EQLTriangle {
  Point vertices[3];
  double colors[3];
  double normals[3];

  public:
    EQLTriangle() {};
    EQLTriangle(double, double, double);
    void print();
};

EQLTriangle::EQLTriangle(double side, double x, double y) {
  vertices[0] = Point(x, y, 0);
  vertices[1] = Point(x + side, y, 0);
  vertices[2] = Point(x + (side / 2.0), y + (sqrt(3.0) * side / 2.0), 0);

  colors[0] = 0.0;
  colors[1] = 1.0;
  colors[2] = 0.0;

  normals[0] = ((vertices[1].y - vertices[0].y) * (vertices[2].z - vertices[0].z)) - ((vertices[1].z - vertices[0].z) * (vertices[2].y - vertices[0].y));
  normals[1] = ((vertices[1].z - vertices[0].z) * (vertices[2].x - vertices[0].x)) - ((vertices[1].x - vertices[0].x) * (vertices[2].z - vertices[0].z));
  normals[2] = ((vertices[1].x - vertices[0].x) * (vertices[2].y - vertices[0].y)) - ((vertices[1].y - vertices[0].y) * (vertices[2].x - vertices[0].x));
};

void EQLTriangle::print() {
  int vertices_size = * ( & vertices + 1) - vertices;
  for (int i = 0; i < vertices_size; i++) {
    printf("Vertices: \n");
    printf("\t[ x: %f, y: %f, z: %f ]\n", vertices[i].x, vertices[i].y, vertices[i].z);
  }

  int colors_size = * ( & colors + 1) - colors;
  printf("Colors: \n");
  printf("\t[ ");
  for (int i = 0; i < colors_size; i++) {
    printf("%f ", colors[i]);
  }
  printf("]\n");

  int normals_size = * ( & normals + 1) - normals;
  printf("Normals: \n");
  printf("\t[ ");
  for (int i = 0; i < normals_size; i++) {
    printf("%f ", normals[i]);
  }
  printf("]\n");
}

int main() {
  srand(time(NULL));

  vector < EQLTriangle > Triangles(3);
  Triangles[0] = EQLTriangle(rand() % 5 + 1, rand() % 500, rand() % 500);
  Triangles[1] = EQLTriangle(rand() % 5 + 1, rand() % 500, rand() % 500);
  Triangles[2] = EQLTriangle(rand() % 5 + 1, rand() % 500, rand() % 500);

  for (size_t i = 0; i < Triangles.size(); ++i) {
    Triangles[i].print();
    printf("\n");
  }
}
