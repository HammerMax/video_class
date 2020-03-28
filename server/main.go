package main

import (
	"fmt"
	"server/modules/video"
)

func main() {
	video.InitVideo()

	as := []video.Attr {
		{
			Vid: "0001",
			Title: "知否知否",
			Image: "url",
		},
		{
			Vid: "0002",
			Title: "test1",
			Image: "url",
		},
		{
			Vid: "0003",
			Title: "庆余年",
			Image: "url",
		},
		{
			Vid: "0004",
			Title: "武林外传",
			Image: "url",
		},
		{
			Vid: "0005",
			Title: "sssssss",
			Image: "url",
		},
	}

	//for _, a := range as {
	//	err := video.Create(&a)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	fmt.Println(as)

	var tmp int
	tmp = 123
	video.Update("0002", map[string]interface{}{"cat": tmp})

	list, err := video.List()
	if err != nil {
		panic(err)
	}

	fmt.Println(list)
}
