package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	//编码: 结构体-->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error : ", err)
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码 jsonStr-->结构体
	decodeMovie := Movie{}
	err = json.Unmarshal(jsonStr, &decodeMovie)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
	}
	fmt.Printf("%s ==>> %v\n", decodeMovie, decodeMovie)
}
