package main

import (
	"fmt"

	"solapi-go"
)

func main() {
	client := solapi.NewClient()

	// 검색조건값
	params := make(map[string]string)
	params["limit"] = "1"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.GetGroupList(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
