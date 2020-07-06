package main

import (
	"context"
	"time"
)

/*
Tạo 1 đối tượng context với 1 value là timestamp hiện tại tính theo ns chạy qua 1 hàm như sau.
 Sau 3s in ra hiệu của thời gian của thời gian hiện tại với biến dữ liệu trong context. in ra màn hình kết quả.
*/
func x(ctx context.Context, k string) int64 {
	var now int64
	select {
	case <-time.After(3 * time.Second):
		now = time.Now().UnixNano()
	}
	return now - (ctx.Value(k)).(int64)
}
