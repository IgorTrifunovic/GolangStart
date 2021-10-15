package main 

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("******************************************")
	go count("ssss")
	go count("aaaaa")
	go count("fffff")
	go count("ddddd")
	go count("==========================================")
	count("fish")
}


func count(thing string){
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 1)
	}
}

