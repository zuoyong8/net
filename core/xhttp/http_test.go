package xhttp

import (
	"fmt"
	"testing"
)

func TestHttp(t *testing.T) {
	// url := "https://wwww.baidu.com"
	// s := joinQuery(url, map[string]interface{}{
	// 	"t": 12223333,
	// 	"s": "b",
	// })

	// h := map[string]string{
	// 	"Content-Type": "application/json",
	// 	"Language":     "zh-CN",
	// }
	// q := map[string]interface{}{
	// 	"phone":    "15221552956",
	// 	"areaCode": "0086",
	// 	"areaFlag": "12333",
	// 	"smsCode":  "6669",
	// 	"device":   2,
	// }
	// bs, err := Post("https://api.wikifxpay.com/api/v1/user/quick/login", h, q)
	// fmt.Println(err, string(bs))

	s, a := JsonToString("http://rating.fx696.com/eahost/ext/eabrokers.json", "brokerid")
	fmt.Println(s, a)

	// var mm map[string]int = make(map[string]int, 0)
	// for k, v := range s {
	// 	mm[v] = k + 1
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(mm["3351410785"])
}
