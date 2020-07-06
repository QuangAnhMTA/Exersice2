package main

import (
	"context"
	"fmt"
	"time"
)
// Thực hiện 1 chương trình với 1 vòng lặp for và 3 lần sleep mỗi lần sleep 3sec Nhưng sau 3s thì kết thúc hàm đấy Sử dụng và tìm hiểu context.
// Nêu được tác dụng của context trong chương trình.
func dungham(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Time Out")
			cancel()
			return ctx.Err()
		default:
			time.Sleep(3 * time.Second)
			println("T:", i)

		}
	}
	fmt.Println("ALL DONE")
	return nil
}
