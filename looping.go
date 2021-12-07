// Looping : control flow 중 하나
package main

import (
	"fmt"
)

// Agenda

// For statements
// - simple loops
// - exiting early
// - looping through collections

//basic for loop

// func main() {
// 	for i := 0; i < 5; i = i+2 { //for i:= 0 -> 여기를 initializer라고 함
// 		fmt.Println(i)			// i < 5 -> generator
// 	}							// i++, i=i+2 -> incrementor
// }

// func main() {
// 	for i, j := 0, 0; i < 5; i, j = i+1, j+1 { //i,j = i++, j++ 이렇게 쓰면 안 돼!!
// 		fmt.Println(i)
// 	}
// }

// increment operation은 statement이고 expression이 아니다..!

//https://shoark7.github.io/programming/knowledge/expression-vs-statement
// expression(수식) : 하나 이상의 값으로 표현(reduce)될 수 있는 코드.
// expression은 evaluate가 가능해서 하나의 값으로 환원됨
// expression은 statement의 부분 집합

// statement: 실행 가능한 최소의 독립적인 코드 조각.
// 프로그래밍하면서 컴파일러가 이해하고 실행할 수 있는 모든 구문은 statement
// statement는 흔히 한 개 이상의 expression과 프로그래밍 키워드를 포함하는 경우가 많음

// 어떤 statement는 expression이지 않다. `return 3` 이런 구문은 함수에서 3을 반환한다는 의미일 뿐이지 평가해서 3이 나오지 않는다.

// 다른 예
// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		if i % 2 == 0 {
// 			i /= 2
// 		} else {
// 			i = 2*i + 1
// 		}
// 	}
// }

// 간단하게 쓰기
// func main() {
// 	// 1) 이렇게 써도 됨
// 	i := 0 // i는 main function에 속해있음
// 	for ; i < 5; i++ { //i를 여기에 쓰면(i := 0) for loop에 속해있음
// 		fmt.Println(i)
// 	}

// 	// 2) 이렇게 써도 됨2
// 	j := 0;
// 	for j < 5 {
// 		fmt.Println(j)
// 		j++
// 	}
// }

//Do while loop 구현. break로 멈춤
// func main() {
// 	i :=0
// 	for {
// 		fmt.Println(i)
// 		i++
// 		if i == 5 {
// 			break
// 		}
// 	}
// }

// continue를 이용해서 홀수만 출력
// continue는 for loop에서 일단 나온 다음 다음 i로 넘어가게 함
// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

// Nested Loop와 Label
// func main() {
// Loop:
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 3; j++ {
// 			fmt.Println(i*j)
// 			if i*j >= 3 {
// 				//break // 안쪽 loop만 벗어남
// 				break Loop //바깥쪽 loop을 벗어남
// 			}
// 		}
// 	}
// }

// func main() {
// 	s := []int{1,2,3}
// 	// for range 몇 번이나 루프 돌지 모를 때
// 	for k, v := range s{ //index, value
// 		fmt.Println(k,v)
// 	}
// }

// func main() {
// 	statePopulations := map[string]int{
// 		"California" : 293830984,
// 		"Texas" : 3434343402,
// 		"Florida" : 3442032933,
// 		"New York" : 193923023,
// 		"Pennsylvania" : 1239343434,
// 		"Illinois" : 128032032,
// 		"Ohio" : 11614934,
// 	}
// 	for k, v := range statePopulations {
// 		fmt.Println(k, v)
// 	}
// }

func main() {
	s := "hello Go"
	for k, v := range s {
		fmt.Println(k,string(v))
	}
}

// Channel

// Summary
// For statements
// - simple loops
// -- for initializer; test; incrementer {}
// -- for test {}
// -- for {}

// exiting early
// - break
// - continue
// - labels

// Looping over collections
// - arrays, slices, maps, strings, channels
// - for k,v := range collection {}
