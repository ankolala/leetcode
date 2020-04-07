package main

import (
	"strings"
	"math"
	"fmt"
	"time"
)

func myAtoi(str string) int {
	return convert(clean(str))
}

func clean(s string) (sign int, abs string) {
	// 先去除首尾空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	// 判断第一个字符
	switch s[0] {
	// 有效的
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
		// 有效的，正号
	case '+':
		sign, abs = 1, s[1:]
		// 有效的，负号
	case '-':
		sign, abs = -1, s[1:]
		// 无效的，当空字符处理，并且直接返回
	default:
		abs = ""
		return
	}
	for i, b := range abs {
		// 遍历第一波处理过的字符，如果直到第i个位置有效，那就取s[:i]，从头到这个有效的字符，剩下的就不管了，也就是break掉
		// 比如 s=123abc，那么就取123，也就是s[:3]
		if b < '0' || '9' < b {
			abs = abs[:i]
			// 一定要break，因为后面的就没用了
			break
		}
	}
	return
}

// 接收的输入是已经处理过的纯数字
func convert(sign int, absStr string) int {
	fmt.Println("sign",sign,"absStr",absStr)
	absNum := 0
	for _, b := range absStr {
		// b - '0' ==> 得到这个字符类型的数字的真实数值的绝对值
		absNum = absNum*10 + int(b-'0')
		// 检查溢出
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
			// 这里和正数不一样的是，必须和负号相乘，也就是变成负数，否则永远走不到里面
		case sign == -1 && absNum*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * absNum
}

func pickupDist(waypoints []int, dists []int) int{
	// result 拼车行程累计接驾距离、driverPos 司机索引、order_size拼车订单数
	result, driverPos, order_size := 0, 0, len(waypoints)/2

	for i:=0;i<len(waypoints);i++ {
		if waypoints[i] == -1 {
			driverPos = i
		}
	}

	for i:=0;i<len(waypoints);i++ {
		if waypoints[i] % 2 == 0 && driverPos < i{
			for j:=driverPos;j<i;j++{
				result += dists[j]
			}
		}
	}
	// 订单平均接驾距离
	return result/order_size
}

func GetNowSubTime(minuend time.Time) time.Duration {
	return time.Now().Sub(minuend)
}

func main() {
	//fmt.Println(clean("-1u3abc-56"))
	//fmt.Println(myAtoi("-123uuu"))
	waypoints := []int{0,2,-1,3,1}
	dists := []int{8870,2897,1307,1306}
	fmt.Println(pickupDist(waypoints, dists))

	now := time.Now()
	for i:=0;i<1000;i++ {
		fmt.Println("123")
	}
	costTime := GetNowSubTime(now)
	fmt.Println(costTime)




}
