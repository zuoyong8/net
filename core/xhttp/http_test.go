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

	q := map[string]interface{}{
		"symbol": "BTCUSDT",
		"limit":  60,
	}

	s, _ := Get("https://q.wikibit.us/api/v3/trades", nil, q)
	fmt.Println(string(s))
}
