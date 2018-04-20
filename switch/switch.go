package main

import (
	"fmt"
	"runtime"
	"time"
)

// switch case

func main(){
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	// 判断距离星期天 还有多久
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	// switch的条件从上到下的执行，当匹配成功的时候停止
	case today + 0 :
		fmt.Println("To day")
	case today + 1 :
		fmt.Println("Tomorrow")
	case today + 2 :
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}


	// ex 3
	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
