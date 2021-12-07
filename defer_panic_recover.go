// Defer, Panic, and Recover -- Control Flow

// Agenda
// - Defer
// - Panic
// - Recover

package main

import (
	"fmt"
	"log"
)

// defer : 함수 마지막에 return 하기 바로 전에 실행
// LIFO order 따르기 때문에 모든 statement에 다 넣으면 마지막에 넣은 순서대로 실행
// func main() {
// 	fmt.Println("start")
// 	defer fmt.Println("middle")
// 	fmt.Println("end")
// }

// func main() {
// 	res, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close() // opening of the rsc와 closing of the rsc를 바로 옆에 쓸 수 있음
// 	robots, err := ioutil.ReadAll(res.Body)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", robots)
// }

// func main() {
// 	a := "start"
// 	defer fmt.Println(a)
// 	a = "end"
// } //start 출력됨

// Panic
// Go에는 exception이 없음...?
// 우리의 application이 뭘 할지 모를 때 Panic함

// func main() {
// 	a, b := 1, 0
// 	ans := a/b
// 	fmt.Println(ans)
// 	}

// Go가 알아서 runtime error를 냄
// panic: runtime error: integer divide by zero

// goroutine 1 [running]:
// main.main()
//         /Users/boyoon/golang/youtube_golang/defer_panic_recover.go:48 +0x20
// exit status 2

// func main() {
// 	fmt.Println("start")
// 	panic("something bad happend")
// 	fmt.Println("end")
// }

// import "net/http"

// func main() { //func(w http...) 이거는 url이 맞을 때마다 호출되는 callback임
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
// 		w.Write([]byte("Hello Go!"))
// 	})
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
//tcp 포트 중복되면 panic

// func main() {
// 	fmt.Println("start")
// 	defer fmt.Println("this was deferred")
// 	panic("something bad happened")
// 	fmt.Println("end")
// }
// 실행순서
// start
// this was deferred // panic 전에 defer함. 즉 application이 패닉하기 전에 defer statement이 리소스를 close할 수 있음
// panic: something bad happened

// func main() {
// 	fmt.Println("start")
// 	defer func() { //annonymous function
// 		if err := recover(); err != nil { //panic하면
// 			log.Println("Error:", err) // log를 print
// 		}
// 	}() //function call을 통해서 이 function을 invoke함
// 	panic("something bad happened")
// 	fmt.Println("end")
// }

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end") //panic후 recover함
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		//	panic(err) //repanic! 이렇게 하면 'end' 출력 안 됨
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

// 출력 순서
// start
// about to panic
// 2021/12/07 22:37:30 Error: something bad happened
// end

// Summary

// Defer
// - Used to delay execution of a statement until function exits
// - Useful to group "open" and "Close" functions together
	// - Be careful in loops
// (Defers) Run in LIFO(last-in, first-out) order
// Arguments evaluated at time defer is executed, not at time of called function execution

// Go에서 뭔가 잘못됐을 떄 error를 반환하거나 panic할 수 있음
// Panic
// panic은 recover할 수 없는 문제가 발생했을 때 사용
// Occur when program cannot continue at all
// - Don't use when file can't be opened, unless it is critical
// - Use for unrecoverable events - cannot obtain TCP port for web server
// Function will stop executing
// - Deferred functions will still fire
// If nothing handles panic, program will exit

// Recover
// Used to recover from panics
// Only useful in deferred functions (application이 패닉하기 시작하면 더이상 다른 함수를 execute하지 못하지만 defer function은 execute함)
// Current function will not attempt to continue, but higher functions in call stack will