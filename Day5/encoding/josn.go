package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var jsonMessage = `[{
	"Name": "Saura",
	"Age": 34,
	"Address": "Solapue"
},{
	"Name": "Gaura",
	"Age": 26,
	"Address": "Solapur"
}]`

type Info struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Addrees"`
}

func main() {
	p := []Info{}

	p2 := make([]map[string]interface{}, 0)
	erer2 := json.Unmarshal([]byte(jsonMessage), &p2)
	fmt.Println(erer2)
	for k, v := range p2 {
		var p4 interface{}
		erer2 = json.Unmarshal(, p4)
		fmt.Println(k, v)
	}

	//fmt.Println(string(p2), erer2)
	err := json.Unmarshal([]byte(jsonMessage), &p)

	if err != nil {
		fmt.Println(err)
	}

	p1 := Info{Name: "Gyncho", Age: 34, Address: "tatychi wadi ghy kela nagar zattwadi"}

	//res, _ := json.Marshal(&p1)

	res1, err := json.MarshalIndent(p1, "", "\t")

	fmt.Println("eier", res1)

	fmt.Println(string(res1))

	fmt.Println(p)
}
