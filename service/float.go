package service

import (
	"fmt"
	"strconv"
)

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}

// MultiHundred : 傳入的數字乘 100 後回傳
func MultiHundred(a float64) float64 {
	qq := a * 100
	qq = FloatRound(qq, 2)
	return qq
}
