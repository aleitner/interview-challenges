#### Company
    Q-Gamea
#### Position
    Software Engineer
### Question
* The answers to these questions should be compilable on GCC, or any ISO
standard C++ compiler (1 file per answer.)

##############################################################
(1)

Please implement the following mathematical function in C++, in the way you
think is best (explain your choice.)

F(0) = 0
F(n) = n + F(n-1)
n=1..20

##############################################################
(2)

Please make a template class called Animal that takes a traits class like
those below, and allows the function CompareAnimals to compile and run
successfully.

#include <iostream>

struct frog_traits
{
	static const char name[];
	static const int number_of_legs = 4;
	static const int number_of_arms = 0;
	static const int weight = 250;		// grams
};

struct zebra_traits
{
	static const char name[];
	static const int number_of_legs = 4;
	static const int number_of_arms = 0;
	static const int weight = 10000;		// grams
};

const char frog_traits::name[] = "frog";
const char zebra_traits::name[] = "zebra";


// Please make the Animal template class here.
...

// End of template class

typedef Animal< frog_traits > Frog;
typedef Animal< zebra_traits > Zebra;


template< class T1, class T2 >
void CompareAnimals( const Animal< T1 >& animal1, const Animal< T2 >&
animal2 )
{
	std::cout << "The difference between a " << animal1.getName() << " and a "
<< animal2.getName() << "\n";

	if ( animal1.numLegs() != animal2.numLegs() )
	{
		std::cout << "is that a " << animal1.getName() << " has " <<
animal1.numLegs() << " legs ";
		std::cout << "and a "     << animal2.getName() << " has " <<
animal2.numLegs() << " legs ";
		std::cout << "\n";
	}

	if ( animal1.numArms() != animal2.numArms() )
	{
		std::cout << "is that a " << animal1.getName() << " has " <<
animal1.numArms() << " arms ";
		std::cout << "is that a " << animal2.getName() << " has " <<
animal2.numArms() << " arms ";
		std::cout << "\n";
	}

	std::cout << animal1.getName() << " is ";
	std::cout << ( (animal1.getWeight() > animal2.getWeight() ) ? "heavier" :
"lighter" );
	std::cout << " than a " << animal2.getName();
	std::cout << "\n";
}

int main( int argc, char *argv[] )
{
	Frog my_frog;
	Zebra my_zebra;

	CompareAnimals( my_frog, my_zebra );

	return 0;
}

##############################################################
(3)

Please use the STL vector class to make a triangle list class, for use in 3d
drawing. For each triangle please store 3 vertex positions, 3 colours, and 1
face normal.



##############################################################
(4)

Please write a program that displays times tables in C++.

##############################################################
(5)

Please fill in the rest of the class "slow_string" class written below.Try
to replicate the STL string class's
functionality as much as possible. This question is to test how good you are
at encapsulating the C standard library functions such as strXXX.

class slow_string
{
	char *data;
public:
	...
}

##############################################################
(6a)

Please make a class (or set of classes) that encapsulate the the standard
functionality required from a 3 element vector. The design should be object
oriented and protect its data members from public access. Please support
simple operations on the vector such as addition, subtraction,
multiplication and division as well as dot and cross product.
You can use/extend some of what you did for (3).


(6b)

Assuming that you had to optimise your vector class for operation on a
modern SIMD processor, explain how you would go about it and what issues you
would have to consider.How would you deal with the issue of supporting both
SIMD and non SIMD architectures with the same class?



##############################################################
(7a)

Given that you have a 2 dimensional array of height values (say 10x10) that
represent a height field.Assume that the distance between each height point
in X and Y is 1 and that from the grid we form a triangulated mesh such that
for each set of 4 grid points A=(x,y), B=(x+1,y) C=(x+1,y+1), and D=(x,y+1)
there are two triangles ABC and ACD (so they are seperated by the line x=y
in the local coordinate system of the grid). Explain how you would find the
*exact* height of an arbitrary point x,y on this mesh. Feel free to reuse
the class you made for (3).

(7b)

Please implement your method in C++.

##############################################################
(8)

If you have this template class

  template <class T, class U>
    struct Typelist
    {
       typedef T Head;
       typedef U Tail;
    };

And this empty class

class NullType
{
};

typedef Typelist<int,Typelist<float,Typelist<char*,NullType> > >
UsableTypes1;
typedef Typelist<short,Typelist<std::string,NullType> > UsableTypes2;


Can you write a template class that takes arbitraty types constructed from
them (as above in UsableTypes1 and UsableTypes2) and as a template argument
and produces a class that contains a number of members of the following
form:

void print(Type _Value) { std::cout << _Value; };

So for UsableTypes1 the resultant class would have members for printing
ints, floats and char*'s, and for UsableTypes2 the resultant class would
have members that print shorts and std::strings.


So for example


MyClass<UsableTypes1> hogehoge;
MyClass<UsableTypes2> blhablha;

void TestMyClass()
{
	int int_val = 0;
	float flt_val = 0.1f;
	const char* char_val = "Hi";
	short short_val = 10;
	std::string str_val = "Hello";

	hogehoge.print( int_val ); // OK
	hogehoge.print( flt_val ); // OK
	hogehoge.print( char_val ); // OK
//	hogehoge.print( short_val); // compile error
//	hogehoge.print( str_val ); // compile error

//	blhablha.print( int_val ); // compile error
//	blhablha.print( flt_val ); // compile error
//	blhablha.print( char_val ); // compile error
	blhablha.print( short_val ); // OK
	blhablha.print( str_val ); // OK
}

This code should work ( and the comment out lines should fail to compile if
you uncomment them).


##############################################################
(9a)

Please think of an algorithm that given 4 arbitrary points produces a smooth
curve that connects all of them.Then write a function in C++ that takes 5
parameters ( 4 points and a "time" parameter between 0 and 1) and returns
the point on the curve at an arbitrary "time".

(9b)

Now write a C++ function that given an arbitrary number of points smoothly
interpolates between them, making use of the function/functions you wrote in
9a if appropriate.
Some form of visual output is a plus.


##############################################################
(10)

Assuming that you have a 2D game that can display hundreds of non-rotating
sprites. Explain how you could efficiently optimize collision detection on
these sprites without having to test every sprite against every other
sprite.

##############################################################

(11)
Think of a question that you feel should be on this test and answer it.
(please don't skip that question)

Date: 06/23/2018