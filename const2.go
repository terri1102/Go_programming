package main

import (
	"fmt"
)

const (
	_ = iota //이런식으로 쓰는 이유는 아래의 const들에 값이 할당되었는지 확인하려고. 할당 안 되어 있을 때 false 뜸. 이때 여기에 사용하는 변수명은 안 중요하니까 _로 씀
	//errorSpecialist = iota //이런식으로 쓰는 이유는 아래의 const들에 값이 할당되었는지 확인하려고. 할당 안 되어 있을 때 false 뜸
	catSpecialist  //iota의 기본 자료형은 int
	dogSpecialist
	snakeSpecialist
)

func main(){
	var specialistType int = catSpecialist
	//var specialistType int //false
	fmt.Printf("%v\n", specialistType == catSpecialist) //true
	fmt.Printf("%v\n", specialistType == dogSpecialist) //false
}