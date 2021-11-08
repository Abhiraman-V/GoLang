// Go program to illustrate the
// use of arithmetic operators
package main

import "fmt"

/*
	   Operators are used to perform operations on variables and values.

	    Arithmetic Operators

			+	Addition
			-	Subtraction
			*	Multiplication
			/	Division
			%	Modulus
			++	Increment
			--	Decrement

		Assignment Operators

			=		Simple Assignment
			+=		Add Assignment
			-=		Subtract Assignment
			*=	   Multiply Assignment
			/=	   Division Assignment
			%=		Modulus Assignment
			&=		Bitwise AND Assignment
			|=	   Bitwise Exclusive OR
			^=	   Bitwise Inclusive OR
			>>=   Right shift assignment
			<<=   Left shift assignment

		Comparison Operators

			==	Equal to
			!=	Not equal
			>	Greater than
			<	Less than
			>=	Greater than or equal to
			<=	Less than or equal to

		Logical Operators

			&& 	Logical and	Returns true if both statements are true
			|| 	Logical or	Returns true if one of the statements is true
			!	   Logical not	Reverse the result, returns false if the result is true

		Bitwise Operators

			& 	AND	                  Sets each bit to 1 if both bits are 1
			|	OR	                     Sets each bit to 1 if one of two bits is 1
			^	XOR	                  Sets each bit to 1 if only one of two bits is 1
			<<	Zero fill left shift	   Shift left by pushing zeros in from the right
			>>	Signed right shift	   Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off

*/
func main() {

	p := 75
	q := 25

	//Arithmetic Operators
	add := p + q
	sub := p - q
	mul := p * q
	div := p / q
	mod := p % q
	fmt.Printf("P + Q = %d", add)
	fmt.Printf("\nP - Q = %d", sub)
	fmt.Printf("\nP * Q = %d", mul)
	fmt.Printf("\nP / Q = %d", div)
	fmt.Printf("\nP %% Q = %d\n", mod)

	//Assignment Operators

	var r = 5

	fmt.Printf("\nR = %d", r)
	var s = r
	fmt.Printf("\nS = %d", s)
	s += r
	fmt.Printf("\nS += R == %d", s)
	s -= r
	fmt.Printf("\nS -= R == %d", s)
	s *= r
	fmt.Printf("\nS *= R == %d", s)
	s /= r
	fmt.Printf("\nS /= R == %d", s)
	s %= 784
	fmt.Printf("\nS (mod)= R == %d", s)
	t := &r
	fmt.Printf("\nT &= R == %d", t)
	s |= 200
	fmt.Printf("\nS |= R == %d", s)
	s ^= 200
	fmt.Printf("\nS ^= R == %d", s)
	s = 5000
	s >>= 3
	fmt.Printf("\nS >>= R == %d", s)
	s <<= 3
	fmt.Printf("\nS <<= R == %d\n", s)

	// Comparison Operators

	var u = 50
	var v = 70

	fmt.Printf("\nU == V == %t", u == v)
	fmt.Printf("\nU != V == %t", u != v)
	fmt.Printf("\nU > V == %t", u > v)
	fmt.Printf("\nU < V == %t", u < v)
	fmt.Printf("\nU >= V == %t", u >= v)
	fmt.Printf("\nU <= V == %t\n", u <= v)

	// Logical Operators

	if u != v && u <= v {
		fmt.Println("\nu!=v && u<=v -> True")
	}

	if u != v || u <= v {
		fmt.Println("u!=v || u<=v -> True")
	}

	if !(u == v) {
		fmt.Println("!(u==v) -> True")
	}

	// Bitwise Operators

	fmt.Printf("\nU & V == %d", u&v)
	fmt.Printf("\nU | V == %d", u|v)
	fmt.Printf("\nU ^ V == %d", u^v)
	fmt.Printf("\nU << V == %d", u<<v)
	fmt.Printf("\nU >> V == %d", u>>v)
}
