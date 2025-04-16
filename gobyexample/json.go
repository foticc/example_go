package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
	Page int
	Data []int
	f    string
}

type Response2 struct {
	Page int   `json:"p"`
	Data []int `json:"d"`
}

func main() {
	bb, _ := json.Marshal(true)
	fmt.Println(string(bb))

	bi, _ := json.Marshal(1)
	fmt.Println(string(bi))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	rep := Response{
		Page: 1,
		Data: []int{1, 2, 3},
		f:    "te",
	}

	repB, _ := json.Marshal(rep)
	fmt.Println(string(repB))

	rep2 := Response2{
		Page: 1,
		Data: []int{1, 2, 3},
	}

	repB2, _ := json.Marshal(rep2)
	fmt.Println(string(repB2))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data["num"].(float64))
	fmt.Println(data["strs"].([]interface{}))
	str := data["strs"].([]interface{})
	fmt.Println(str[0].(string))

	d := `{"Page":1,"Data":[1,2,3]}`
	res := Response{}
	if err := json.Unmarshal([]byte(d), &res); err != nil {
		panic(err)
	}
	fmt.Println(res)

	enc := json.NewEncoder(os.Stdout)
	m := map[string]string{"name": "Alice", "age": "18"}
	enc.Encode(m)
}
