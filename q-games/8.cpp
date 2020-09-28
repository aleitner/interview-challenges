/*
Question:

Can you write a template class that takes arbitraty types constructed from
them (as above in UsableTypes1 and UsableTypes2) and as a template argument
and produces a class that contains a number of members of the following
form:

void print(Type _Value) { std::cout << _Value; };

So for UsableTypes1 the resultant class would have members for printing
ints, floats and char*'s, and for UsableTypes2 the resultant class would
have members that print shorts and std::strings.

SOLUTION: See Below

This question was a little difficult for me.
I ended up writing Contains struct for checking if the Type passed in is inside of the Typelist being used
There is no compilation error however unaccepted data types won't print.

*/

#include <iostream>
#include <string>

class NullType {};

template < class T, class U >
struct Typelist {
  typedef T Head;
  typedef U Tail;
};

typedef Typelist < int, Typelist < float, Typelist <
  const char * , NullType > > >
    UsableTypes1;
typedef Typelist < short, Typelist < std::string, NullType > > UsableTypes2;

// Contains
template < class Type, class Tlist > struct Contains;

template < class T >
  struct Contains < T, NullType > {
    enum {
      value = false
    };
  };

template < class Type, class T, class U >
  struct Contains < Type, Typelist < T, U > > {
    enum {
      value = std::is_same < Type, T > ::value || Contains < Type, U > ::value
    };
  };

template < class T >
  class MyClass {
    public:
      template < typename Type >
      void print(Type _Value) {
        if (Contains < Type, T > ::value == true) {
          std::cout << _Value << "\n";
        }
      };
  };

MyClass < UsableTypes1 > hogehoge;
MyClass < UsableTypes2 > blhablha;

void TestMyClass() {
  int int_val = 0;
  float flt_val = 0.1f;
  const char * char_val = "Hi";
  short short_val = 10;
  std::string str_val = "Hello";

  hogehoge.print(int_val); // OK
  hogehoge.print(flt_val); // OK
  hogehoge.print(char_val); // OK
  //	hogehoge.print( str_val ); // compile error

  //	blhablha.print( int_val ); // compile error
  //	blhablha.print( flt_val ); // compile error
  //	blhablha.print( char_val ); // compile error
  blhablha.print(short_val); // OK
  blhablha.print(str_val); // OK
}

int main() {
  TestMyClass();
}
