package main

import (
	"fmt"
	"time"
)

// Tạo 1 đoạn code sử dụng time.AfterFunc(), sau 100ms thì in ra dòng chữ i'm study

func b9() {

	f := func() {

		fmt.Println("i'm study")
	}
	timer := time.AfterFunc(100*time.Millisecond, f)
	defer timer.Stop()
	time.Sleep(1 * time.Second)
	fmt.Println("Hoàn thành bài tập 2. The end")

}
