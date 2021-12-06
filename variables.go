package main

import (
	"fmt"
	"strconv"
)

//패키지 레벨에서의 변수 선언
// := 사용 불가
var pl int = 42

//블럭으로 선언하기
// 변수들이 다 같은 타입 아니어도 뭔가 관계가 있다는 것 암시
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah lane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	//1) 변수 선언 후 할당
	var i int
	i = 42

	//2) 변수 선언과 할당 동시에
	// 변수의 타입을 확실히 알 때
	var j int = 42
	var p float32 = 0.3

	//3) Go가 변수 타입 유추하게 함
	// 이 변수에 대한 정보가 별로 없을 때 사용
	// 이렇게 해도 int라고 알기는 함.
	// 하지만 c:= 0.3 이런 식으로 쓰면 무조건 float64으로 유추하고 float32로 유추할 수 없음
	// 또한, 함수 밖에서 쓰이는 전역변수는 이렇게 선언할 수 없음!!!!!
	c := 35

	fmt.Printf("%v, %T", j, j)
	fmt.Printf("%v, %T", p, p)
	fmt.Printf("%v", pl)
	fmt.Println(i)

	fmt.Println(c)

}

//Go에서는 선언되기만 하고 사용되지 않은 변수가 있으면 compile error가 남.
//Go의 naming convention
// lowercase: 패키지 레벨. func 안에 있음. 밖으로 드러나지 않음.
// uppercase: 전역적으로 드러남.
// 이름의 길이는 이 변수의 수명을 잘 나타내야 함. 즉 짧게 쓸 변수는 간단하게.
// acronym은 모두 대문자로 쓰기. HTTP, URL, theURL -가독성 위해서

var it int = 35

// 변수 재선언
func redeclare() {
	var t int = 30 
	//t := 12		같은 scope 안에 변수가 이미 선언되어 있기 때문에 다시 선언할 수 없다. 
					// t = 12 이런식으로 재할당 가능
	var it int = 42		//함수 밖에 전역 변수로 it가 이미 선언되어 있지만 inner scope에 있는 변수가 이를 overide(shadowing)함
	

	// type conversion
	var dd float32
	dd = float32(t) //t는 위에서 int였는데 이제 float32로 바뀌었음
	fmt.Println(dd)
	fmt.Println(it)
	//var ad int
	//ad = dd       // 하지만 1.4 같은 float을 int인 변수에 할당하면 compile error남

	//string의 경우
	// Go의 경우 string은 string of bytes
	// j = string(it) 이런식으로 바꾸고 fmt.Printf("%v, %T\n", j,j) 프린트하면 *, string으로 나옴. 유니코드에서 *가 42여서 그럼
	// 따라서 string으로 conversion하고 싶으면 "strconv" 패키지를 불러와서 사용해야 함
	var j string
	j = strconv.Itoa(it)
	fmt.Printf("%s", j)
}


// Variable visibility
// 전역변수로 소문자 변수 선언 : scope to this package
// 전역변수로 대문자 변수 선언 : exported from package and globally visible
// block scope : 함수 안은 하나의 scope

//-----------------------------------------------------
//- Variable declaration
//-- var foo int
//-- var foo int = 42
//-- foo := 42

//- Can't redeclare variables, but can shadow them

//- All variables must be used

//- Visibility
//-- lower case first letter for package scope
//-- upper case first letter to export
//-- no private scope

//- Naming conventions
//-- Pascal or camelCase
//--- capitalize acronyms(HTTP, UTL)
//-- as short as possible
//--- longer names for longer lives

//- Type conversions
//-- destinationType(variable)
//-- use strconv package for strings