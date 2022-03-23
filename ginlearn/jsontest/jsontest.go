package main

import (
	"fmt"
	"os"
	//"encoding/json"
	"github.com/goccy/go-json"
)


type ColorGroup struct {

	ID	int
	Name string
	Colors []string
}

func main() {

	cg := ColorGroup{
		ID: 1,
		Name: "Reds",
		Colors: []string{"Red1", "Red2", "Red3", "Red4"},
	}

	bytes, err := json.Marshal(cg)
	if err != nil{
		fmt.Println("error:", err)
	}

	os.Stdout.Write(bytes)

	cg2 := ColorGroup{}
	err = json.Unmarshal(bytes, &cg2)
	if err != nil{
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v", cg2)
}

