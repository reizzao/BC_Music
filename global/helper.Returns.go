package global

import (
	"encoding/json"
	"fmt"
)

func ReturnJson(d any) string {
	res, _ := json.Marshal(d)
	return string(res)
}

func Console(d any) {
	res, _ := json.Marshal(d)
	fmt.Println(string(res))
}


