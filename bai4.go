package main

import (
	"time"
)
// Tính ra số phút(number_of_minutes) của mốc thời gian sau 1592190294764144364
func numberminutes(n int64) int64 {
	p := n / (60 * int64(time.Second))
	return p
}
