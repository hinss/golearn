package main

func test() int {

	s := make([]int,999,1000)

	return s[88]
}


func main() {

	m := make(map[string]string, 2)
	m["name"] = "hins"

	delete(m, "name")

}
