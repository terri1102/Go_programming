// Collection type : Arrays and Slices
package main

import (
	"fmt"
)

func main() {
	// 비슷한 변수 선언시 이런식으로 쓰지 말고 array로 쓰기
	// grade1 := 97
	// grade2 := 85
	// grade3 := 93

	// Go의 특이점
	// array의 요소들이 메모리에 연속적(contiguous)으로 들어가 있다
	//grades := [3]int{97,85,93} //array사이즈를 두번이나 선언하고 있음
	grades2 := [...]int{98,56,47}
	fmt.Printf("Grades: %v,%v, %v", grades2[0], grades2[1], grades2[2])
	var students [3]string
	fmt.Printf("Students: %v", students)
	students[0] = "Lisa"
	fmt.Printf("students %v", students)
	fmt.Printf("Student #1: %v\n", students[1])

	var identityMatrix [3][3]int = [3][3]int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1}}
	fmt.Println(identityMatrix)

	a := [...]int{1,2,3}
	b := a //b는 a의 모든 요소를 포함한 array의 복사본. array가 클 경우 메모리 차지 및 시간 오래 걸림
//	b := &a //b는 a와 똑같은 데이터를 가리킬 것
	b[1] = 5
	fmt.Println(a) //[1 2 3]
	fmt.Println(b) // &[1 5 3] 
	fmt.Println(a) // [1 2 3] b는 a의 복사본이어서 영향 안 받음

	//slice 문법
	a2 := []int{1,2,3} // 슬라이스
	b2 := &a2
	b3 := a2
	//b2[1] = 5 b는 a의 데이터를 가리키기 때문에 슬라이싱으로 값을 할당할 수 없음
	fmt.Println(a2) //[1 2 3]
	fmt.Println(b2) //&[1 2 3]
	fmt.Println(b3) //[1 2 3]
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))
	aa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //slice : 가변길이 <-> array: 고정길이
	bb := aa[:] //slice of all elements
	cc := aa[3:] //slice from 4th element to end
	dd := aa[:6] // slice first 6 element
	ee := aa[3:6] //slice the 4th, 5th, 6th elements
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(ee)

	ax := []int{}
	fmt.Println(ax)
	fmt.Printf("Length: %v\n", len(ax))
	fmt.Printf("Capacity: %v\n", cap(ax))
	ax = append(ax,1) //built-in method //zero sized array보다 사이즈가 하나 더 큰 array를 새로 만들었음
	fmt.Println(ax)
	fmt.Printf("Length: %v\n", len(ax))
	fmt.Printf("Capacity: %v\n", cap(ax))


	ax = append(ax, 2,3,4,5) 
	//ax = append(ax, []int{2,3,4,5}...) //spread operator: 뒤에 ...붙이면 append(ax, []int{2,3,4,5}...)
	fmt.Println(ax)
	fmt.Printf("Length: %v\n", len(ax)) //5
	fmt.Printf("Capacity: %v\n", cap(ax)) //6

	//slice를 stack으로 취급할 수도 있음
	as := []int{1,2,3,4,5}
	bs := as[1:] //맨 앞 자리 요소 제거
	cs := as[:len(as)-1] //맨 뒷자리 요소 제거
	ds := append(as[:2], as[3:]...) //중간 요소 제거
	fmt.Println(ds) //[1 2 4 5]
	fmt.Println(as) //[1 2 4 5 5] //as도 영향을 받게 됨..! 나중에 for loop에서 어떻게 다뤄야 하는지 나옴

	fmt.Println(as,bs,cs)
}

//Summary

//Array
//- Collection of items with same type
//- Fixed size
//- Declaration styles
//-- a:= [3]int{1,2,3}
//-- a:= [...]int{1,2,3} //나중에 크기 키울 수 있음
//-- var a [3]int
//- Access via zero-based index
//-- a := [3]int {1,3,5} // a[1] == 3
//- len function returns size of array
//- copies refer to different underlying data

//Slices
//- backed by array
//- creation styles
//-- slice existing array or slice
//-- literal style : runtime에 결정됨
//-- via make function
//--- a := make([]int, 10) //create slice with capacity and length == 10
//--- a := make([]int, 10, 100) // slice with length == 10 and capacity == 100
//- len function retuns length of slice
//- cap function rturns length of underlying array
//- append function to add elements to slice
//-- may cause expensive copy operation if underlying array is too small
//- copies refer to same underlying array