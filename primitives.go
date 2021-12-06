//Primitives
//Agenda
//Boolean type
//Numeric types (integers, floating points, complex numbers)
//Text types

// Boolean variables

package main

import (
	"fmt"
)

func main() {
	n := 1 == 1
	m := 1 == 2
	var a bool = false
	fmt.Printf("%v, %T\n", a, a) //false, bool
	fmt.Printf("%v, %T\n", n,n)
	fmt.Printf("%v, %T\n", m,m)

	// Go에서는 모든 변수를 선언할 때 zero value가 할당되기 때문에, 선언만 하고 프린트해도 값(zero value)이 나옴
	var mn bool
	fmt.Printf("%v, %T\n", mn, mn) //false, bool. Boolean의 zero value는 false

	// int
	// int is a signed integer type that is at least 32 bits in size
	// int8(-128~127), int16(-32768~32767), int32(-2147483648~2147483647), int64(-9,223,372,036,854,775,808~9,223,372,036,854,775,807) 가능
	// uint : unsigned integer

	c := 10
	d := 3
	fmt.Println(c+d)
	fmt.Println(c-d)
	fmt.Println(c*d)
	fmt.Println(c/d) // 10을 3으로 나눈 값의 int 값
	fmt.Println(c%d) // 

	var x int = 10
	var y int8 = 3
	//fmt.Println(x+y) // mismathced types int and int8
	fmt.Println(x+int(y))

	// bit operator
	a1 := 10 //1010
	b1 := 3  //0011
	fmt.Println(a1 & b1) //0010
	fmt.Println(a1 | b1) //1011
	fmt.Println(a1 ^ b1) //XOR 연산 //1001
	fmt.Println(a1 &^ b1) //AND NOT 비트연산 //0100

	a2 := 8 //2^3
	fmt.Println(a2 << 3) //2^3 * 2^3 = 2^6
	fmt.Println(a2 >> 3) //2^3/2^3 = 2^0

	//floating point numbers : decimal number를 initialize할 때 float으로 해야함
	//float32, float64
	f1 := 3.14
	f1 = 13.7e72 //13.7 * 10^(72)  부동소숫점 방식으로 써야 함
	//f1 = 13.7e-72 // 13.7 * 10^(-72)
	f1 = 2.1E14 // 2.1 * 10^(14)
	fmt.Printf("%v, %T\n", f1, f1) 
	//가능한 연산
	// 	f1 + f2, f1 - f2, f1/f2, f1 * f2 //bit shift불가, % 불가

	//complex number
	//complex64, complex128
	var c1 complex64 = 1+2i
	fmt.Printf("%v, %T\n", c1, c1)
	//복소수를 실수와 허수로 분리
	var comp complex64 = 1+2i
	fmt.Printf("%v, %T\n", real(comp), real(comp)) //1, float32(complex64여서, complex128이면 float64나옴)
	fmt.Printf("%v, %T\n", imag(comp), imag(comp)) //2, float32
	var comp2 complex128 = complex(5,12)
	fmt.Printf("%v, %T\n", comp2, comp2)

	// Text type
	s1 := "this is a string"
	//s1[2] = "u" string은 이렇게 변경할 수 없음
	fmt.Printf("%v, %T\n", string(s1[2]), s1[2]) //python처럼 array로 취급 가능 s[2]
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s1 + s2, s1+s2) //string끼리 더하는 거 가능, 빼는 건 안 됨

	//slice of bytes
	d1 := []byte(s1) 
	fmt.Printf("%v, %T\n", d1, d1) //[116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
	//전송할 때 byte slice로 작업할 일이 많음

	// Rune type, in single quote
	//utf-32 이지만 alias for int32
	// 언제 사용하냐면 utf-32를 ReadByte로 불러올 때 소실될 정보들이 ReadRune으로 불러오면 데이터가 저장됨
	r := 'a'
	fmt.Printf("%v, %T\n", r,r) //97, int32
}

//Summary

//- Boolean type
//-- values are true of false
//-- not an alias for other types(e.g. int)
//-- zero value is false

//Numeric types
//- Integer
//-- Signed integers
//--- int type has varing size, but min 32 bits
//--- 8 bit(int8) through 64 bit(int64)
//-- Unsigned integers
//--- 8 bit(byte and uint8) through 32 bit(uint32)
//-- Arithmetic operations
//--- Addition, subtraction, multiplication, division, remainder
//-- Bitwise operations
//--- And, or, xor, and not
//-- Zero value is 0
//-- Can't mix types in same family(uint16 + uint32 = error)

//- Floating point numbers
//--- follows IEEE-754 standard
//--- zero value is 0
// 32 and 64 bit versions
//-- Literal styles
//---decimal(3.14)
//---exponential(13e18 or 2E10)
//---mixed(13.7e12)
//-- Arithmetic operations
//--- addition, subtraction, multiplication, division

//- Complex numbers
//--- zero value is (0+0i)
//--- 64 and 128 bit versions
//--- built-in functions
//---- complex - make complex number from two floats
//---- real - get real part as float
//---- imag - get imaginary part as float
//--- arithmetic operations
//---- addition, subtraction, multiplication, division


// Text types
//- Strings
//--- UTF-8
//--- immutable
//--- can be concatenated with plus(+) operator
//--- can be converted to []byte
//- Rune
//--- UTF-32
//--- alias for int32
//--- special methods normally required to precess ex)strings.Reader #ReadRune
