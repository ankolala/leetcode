package main

import (
	"math/rand"
	"time"
	"encoding/json"
	"fmt"
)

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min) + min
	return randNum
}


func appendJson(rawJson string, addKv map[string]interface{}) string {

	rawMap := map[string]interface{}{}

	if err := json.Unmarshal([]byte(rawJson), &rawMap); err != nil {
		return ""
	}

	for k, v := range addKv {
		rawMap[k] = v
	}

	return objToStr(rawMap)
}

func objToStr(data interface{}) string {
	s, _ := json.Marshal(data)
	return string(s)
}

// ConvIDLongToShort 将长ID转换为短ID
func ConvIDLongToShort(longID uint64) (shortID uint64) {
	shortID = longID & 0xffffffffffff
	return
}

func main(){
	fmt.Println(ConvIDLongToShort(281474991375772))
	//fmt.Println(GenerateRangeNum(0, 2))
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(2)) //生成0-1的随机整数
	extendParam := "{\"mrp_strategy\":1,\"fast_rank_type\":0, \"dynamic_route\":8}"

	addKv := map[string]interface{}{"route_strategy":`{"auto_recommend":0}`}


	fmt.Println(appendJson(extendParam, addKv))


}
