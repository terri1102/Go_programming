// Agenda

// Maps
// -what are they?
// -creating
// -manipulation

// Structs
// - what are they?
// - creating
// - naming conventions
// - embedding
// - tags

package main

import (
	"fmt"
	"reflect"
)



// 하나의 Struct 안에 다양한 type 사용 가능
// 아래의 struct은 대문자니까 다른 패키지에서 접근 가능하지만, fieldname은 소문자여서 접근 불가
type Doctor struct {
	number int //fieldname : type associated with the field
	actorName string
	companions []string
}


func main() {
	//make로 map 만들기
	statePopulations := make(map[string]int) //map 선언만
	statePopulations = map[string]int{
		"California" : 39250017,
		"Texas" : 27862596,
		"Florida" : 20612439,
		"New York" : 19745289,
		"Pennsylvania" : 12802503,
		"Illinois" : 12801539,
		"Ohio" : 11614373,
	}

	// 위와 동일한 결과
	// statePopulations := map[string]int{
	// 	"California" : 39250017,
	// 	"Texas" : 27862596,
	// 	"Florida" : 20612439,
	// 	"New York" : 19745289,
	// 	"Pennsylvania" : 12802503,
	// 	"Illinois" : 12801539,
	// 	"Ohio" : 11614373,
	// }
	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "Georgia")
	//근데 statePopulations에서 Georgia를 지워도 print하면 0으로 나오긴 함
	fmt.Println(statePopulations["Georgia"])
	//key에 오타가 있어도 print하면 0으로 나옴 statePopulations["Oho"]
	//이러면 value가 0인지, key가 없는 건지 알 수 없기 때문에 아래와 같이 사용함
	pop, ok := statePopulations["Oho"]
	//_, ok := statePopulations["Oho"] //pop 하지 말고 그냥 key있나 확인하고 싶을 때 사용
	fmt.Println(pop, ok) //0 false
	fmt.Println(statePopulations["Ohio"])
	//map의 길이 구하기
	fmt.Println(len(statePopulations))
	//map을 복사한 새로운 map에서 key를 지우면 원래 map에서도 지워짐
	sp := statePopulations
	delete(sp, "Georgia")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	//-------------------------------------------------------
	aDoctor := Doctor{
		number : 3,
		actorName: "Jon Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	aDoctor2 := struct{name string}{name: "John Pertwee"} //첫번째 {}는 타입 선언, 두번째 {}는 값으로 initialize
	anotherDoctor := aDoctor2 //위의 aDoctor 개체와 독립적. anotherDoctor를 수정해도 aDoctor에 영향을 주지 않음
	anotherDoctor.name = "Tom Baker" //
	sameDoctor := &aDoctor2
	sameDoctor.name = "Tom Baker"
	fmt.Println(sameDoctor) //&{Tom Baker}
	fmt.Println(aDoctor2) //{Tom Baker}
	//fmt.Println(aDoctor) 이렇게 전체 print하는 건 추천하지 않음. field 선언시 순서와 할당시 순서가 다르면..순서 이상하게 인식함. 그래서 field별로 print하는 것이 더 나음
	fmt.Println(aDoctor.actorName) //Jon Pertwee
	fmt.Println(aDoctor.companions[1]) //Jo Grant

	
	//Embedding
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b) //{{Emu Australia} 48 false}


	bb := Bird{
		Animal: Animal{Name:"Emu", Origin: "Australia"},
		SpeedKPH : 48,
		CanFly: false,
	}
	fmt.Println(bb.Name)

	
	cbird := Bird{
		Animal : Animal {Name: "Emu", Origin: "Australia"},
				SpeedKPH :48,
				CanFly: false,
	}
	fmt.Println(cbird.Name)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println("field.Tag")
}



type Animal struct { //얘를 Bird 아래에 둬도 잘 작동함 
	Name string
	Origin string
}

type Bird struct {
	Animal	//Animal struct을 상속 받는 것처럼 구현(상속과 비슷하게 구현되지만 여전히 Animal과 다른 struct임)
	SpeedKPH float32
	CanFly bool
}

//Tag
type Animal struct {
	Name string `required max:"100"` //tag는 백틱키로 감쌈. reflect 패키지를 가져와서 사용
	Origin string
}

//Summary

//Maps
//- Collections of value types that are accessed via keys
//- Created via literals or via `make` function
//- Members accessed via [key] syntax
//-- myMap["key"] = "value"
//- Check for presence with "value, ok" form of result
//- Multiple assignments refer to same underlying data

// Stucts
//- Collections of disparate data types that describe a single concept
//- Keyed by named fields
//- Normally created as types, but anonymous structs are allowed
//- Structs are value types
//- No inheritance, but can use composition via embedding
//- Tags can be added to struct fields to dscribe field
