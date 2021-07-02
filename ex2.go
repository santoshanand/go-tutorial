package main

import (
	"fmt"
	"time"
)

func timeMap(v interface{}) {
	res, ok := v.(map[string]interface{})
	fmt.Println(ok)
	if ok {
		res["update_at"] = time.Now()
	}
}
func main() {
	res := map[string]interface{}{
		"Santosh": 1,
	}

	timeMap(res)
	timeMap(1)

	fmt.Println(res)
}