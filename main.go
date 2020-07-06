package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Bài 1
	dung3(3)
	// Bài 2
	tinhngay()
	// Bài 3
	ctx := context.Background()
	err := dungham(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
	// Bài 4
	p := numberminutes(1592190294764144364)
	fmt.Println(p)
	// Bài 5
	dayweek(1592190385)
	/*
		Trong golang mặc định thì thời gian dạng số được sử dụng với các loại mốc đơn vị là second và nanosecond
	*/
	// Bài 7
	t := time.Now().UnixNano()
	k := "key"
	ctx1 := context.WithValue(context.Background(), k, t)
	println(x(ctx1, k))
	//Bài 8
	intervaltime()
	// Bài 9
	b9()
}
