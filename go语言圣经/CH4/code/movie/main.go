package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	// 没有设置Tag，输出为Title
	Title string
	// 该成员以Tag:releaseed名字输出
	Year int `json:"released"`
	// omitempty，如果该字段为空或者零值，则输出该成员
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	// 使用MarshalIndent可以产生整齐缩进的输出
	//data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println()

	// 只定义了Title，选择性的解码JSON中的Title
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
