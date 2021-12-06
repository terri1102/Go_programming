//Agenda

// If statements
//- Operators
//- If - else and if - else if statements

// Switch statements
//- Simple cases
//- Cases with multiple tests
//- Falling through
//- Type switches

package main

import (
	"fmt"
	"math"
)

func main() {
	// if false {
	// 	fmt.Println("The test is true")
	// }
	
	// Initializer Syntax
	statePopulations := map[string]int{
		"California" : 39250017,
		"Texas" : 27862596,
		"Florida": 206123232,
		"New York" : 323434,
		"Pennsylvania" : 129934903,
		"Illinois": 343040234,
		"Ohio": 34343423,

	}
	//여기 if문 시작하는 부분이 initializer syntax임
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	//숫자 맞추기 게임
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}
	fmt.Println(number<=guess, number >= guess, number != guess)

	// 숫자 맞추기 게임2
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	//not operator !
	fmt.Println(!true)

	//	블럭으로 묶을 수도 있음
	if guess < 1  {
		fmt.Println("The guess must be greater than 1")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100")
	} else {
			if guess < number {
			fmt.Println("Too low")
			}
			if guess > number {
				fmt.Println("Too high")
			}
			if guess == number {
				fmt.Println("You got it")
			}
			fmt.Println(number <= guess, number >= guess, number != guess)
	}
	
	// floating point 비교
	// 근사값을 쓰기 때문에 바로 동등비교하지 말고
	// error value를 정한 다음에 threshold 보다 큰지 체크하는 식으로 비교
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) { //다르다고 나옴
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
	myNum2 := 0.12323
	if math.Abs(myNum2 / math.Pow(math.Sqrt(myNum2),2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}


	// Switch 문
	switch 2 { // tag라고 함
	case 1, 3, 5:
		fmt.Println("one, three, or five")
	case 2, 4, 6:
		fmt.Println("two, four, or siz")
	default:
		fmt.Println("not one or two")
	}


	// tagless syntax
	i := 10
	switch {
	case i <= 10 :
		fmt.Println("less than or equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
	//fallthrough : 논리없이 next case statement이 execute되게 함


	// type
	var e interface{} = [3]int{}
	switch e.(type) {
	case int:
		fmt.Println("e is an int")
		//중간에 case 나가기
		break
		fmt.Println("This is not goint to be run")

	case string:
		fmt.Println("e is string")
	case float64:
		fmt.Println("e is a float64")
	case [3]int:
		fmt.Println("e is [3]int") //array type이 같으려면 같은 데이터타입과 같은 사이즈여야 함
	default:
		fmt.Println("e is another type")
	}
}

//Short circuiting:
// Go is gonna lazily evaluate statements
// or test,and test 패스하면 일부 syntax 건너 뛰기도 함

//Summary

//If statements
// initializer
// comparison operators
// logical operators
// short circuiting
// if - else statements
// if - else if statements
// Equality and floats

// Switch statements
// switching on a tag
// cases with multiple tests
// Initializers
// switches with no tag
// fallthrough
// Type switches
// Breaking out early