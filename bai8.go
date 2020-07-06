package main

import (
	"context"
	"fmt"
	"time"
)

//Tạo 1 interval time sao cho cứ 100ms in ra dùng chữ ${time.Now().Unix()} done
func intervaltime() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for t := range time.Tick(100 * time.Millisecond) {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(t, time.Now().Unix(), "done")
		}
	}
}
