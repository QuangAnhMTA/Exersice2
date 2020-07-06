package main

import (
	"fmt"
	"time"
)
// Tính ra số ngày trong tuần(dạng string và number) của mốc thời gian sau 1592190385
func dayweek(t int64) {

	weekday := time.Unix(0, t).Weekday()
	fmt.Println(weekday)
	fmt.Println(int(weekday))
}
