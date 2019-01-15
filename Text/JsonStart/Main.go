package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	m1 := Movie{
		Title:  "WallE",
		Year:   2008,
		Color:  true,
		Actors: []string{"班·布鲁特", "娜丽莎·奈特", "佛莱德·威拉特"},
	}
	js, err := json.MarshalIndent(m1, " ", "    ")

	if err == nil {
		fmt.Printf("%s\n", js)
	} else {
		fmt.Println(err)
	}

	var mov Movie
	if err := json.Unmarshal(js, &mov); err != nil {
		log.Fatal("JSON unmarshaling failed")
	}
	fmt.Println(mov)

}
