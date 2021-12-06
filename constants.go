//Agenda
//- Naming convention
//- Typed constants
//- Untyped constants
//- Enumerated constants
//- Enumerated expressions

package main

import (
	"fmt"
)

const a int16 =27

//enumerated const
const (
	app = iota
	//bpp = iota
	//cpp = iota
	bpp // 할당해주지 않아도 app의 값을 보고 compiler가 유추함. iota에 바인딩됨. 즉 iota를 맨 위의 상수에 할당하면 순서대로 값이 할당되고 자료형도 첫 줄과 동일해짐
	cpp 
)

const (
	a2 = iota
)



func main() {
	fmt.Printf("%v\n",app) //0
	fmt.Printf("%v\n",bpp) //1
	fmt.Printf("%v\n", cpp) //2




	const a int = 14 // 전역 변수로 선언된 값 뿐 아니라 그 타입(int)도 main func 안의 값이 shadow함
	const myConst int = 42
	//const myconst2 float64 = math.Sin(1.57) // const로 선언하는 값들은 Runtime에 결정되어야 하는 값이면 안 됨. math.Sine()은 Compile time에 execute 되어야 하기 때문에 안 됨
	// primitive type이면 가능함.(int, string, float32, bool) array 등 mutable 해도 const로 선언 불가

	//myConst = 3 const로 선언하면 값을 변경 할 수 없음
	fmt.Printf("%v, %T\n", myConst, myConst)

	// const와 var은 타입만 같으면 연산 가능
	const t int = 42
	var p int = 27
	fmt.Printf("%v, %T\n", t+p, t+p) //69, int

	const ab = 42
	var bb int16 = 27
	fmt.Printf("%v, %T\n", 42 +bb, 42+bb) //69, int16

	
}


//Summary - Constants

// - Immutable, but can be shadowed
// - Replaced by the compiler at compile time
// -- value must be calculable at compile time

// - Named like variables
// -- PascalCase for exported constants
// -- camelCase for internal constants

// - Typed constants wok like immutable variables
// -- can interoperate only with same type

// - Untyped constants work like literals
// -- can inteoperate with similar types

// - Enumerated constants
// -- special symbol iota allows related constants to be created easily
// -- Iota starts at 0 in each const block and increments by one
// -- watch out of constant values that match zero values for variables

// - Enumerated expressions
// -- Operations that can be determined at compile time are allowed
// --- Arithmetic
// --- Bitwise operations
// --- Bitshifting
