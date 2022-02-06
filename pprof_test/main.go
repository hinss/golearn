package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"
)

var datas []string

func main() {

	go func() {
		for {

			log.Printf("len: %d", Add("go-programming-tour-book"))
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}
