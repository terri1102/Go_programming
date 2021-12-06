package main

import (
	"fmt"
)

const (
	_ = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) 
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	//변수들마다 한 바이트씩 나타나게 됨
	isAdmin = 1 << iota //1
	isHeadquarters	//2
	canSeeFinancials //4

	canSeeAfrica //8
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)


func main() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB) //3.73GB
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope //100101 //이렇게 하면 세 가지 const의 값을 아주 효과적으로 저장할 수 있음.
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin) // bit masking //true
	fmt.Printf("IS HQ? %v\n", isHeadquarters & roles == isHeadquarters) //bitmasking //true

}

