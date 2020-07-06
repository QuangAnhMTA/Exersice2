package main

import (
	"fmt"
	"time"
)
// Việt 1 đoạn chương trình tính ra ngày hiện tại theo timestamp. Điểm mốc từ mức 0 tại 1/1/1970
func tinhngay() {
	sec := time.Now().Unix()
	fmt.Println("So ngay hien tai theo cot moc 1/1/1970 la:", sec/(60*60*24))
}
