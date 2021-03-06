package main

import (
"fmt"
"time"
)

var slaveDns = map[int]map[string]interface{}{
	0: {"connectstring": "root@tcp(172.16.0.164:3306)/shiqu_tools?charset=utf8", "weight": 2},
	1: {"connectstring": "root@tcp(172.16.0.165:3306)/shiqu_tools?charset=utf8", "weight": 4},
	2: {"connectstring": "root@tcp(172.16.0.166:3306)/shiqu_tools?charset=utf8", "weight": 8},
}

var i int = -1 //表示上一次选择的服务器
var cw int = 0 //表示当前调度的权值
var gcd int = 2 //当前所有权重的最大公约数 比如 2，4，8 的最大公约数为：2

func getDns() string {
	for {

		i = (i + 1) % len(slaveDns) // slaveDns size always equals to 3
		fmt.Println(i)
		if i == 0 {
			cw = cw - gcd
			if cw <= 0 {
				cw = getMaxWeight() //MaxWeight equals to 8
				if cw == 0 {
					return ""
				}
			}
		}
		fmt.Println("cur choice:",slaveDns[i]["connectstring"].(string), slaveDns[i]["weight"].(int),cw)
		if weight, _ := slaveDns[i]["weight"].(int); weight >= cw {
			return slaveDns[i]["connectstring"].(string)
		}
	}
}

// genMaxWeight函数用于获取最大权值
func getMaxWeight() int {
	max := 0
	for _, v := range slaveDns {
		if weight, _ := v["weight"].(int); weight >= max {
			max = weight

		}
	}
	return max

}

func main() {

	note := map[string]int{}

	s_time := time.Now().Unix()

	for i := 0; i < 10; i++ {
		s := getDns()
		fmt.Println("return:",s)
		if note[s] != 0 {
			note[s]++
		} else {
			note[s] = 1
		}
	}

	e_time := time.Now().Unix()

	fmt.Println("total time: ", e_time-s_time)

	fmt.Println("--------------------------------------------------")

	for k, v := range note {
		fmt.Println(k, " ", v)
	}
}
