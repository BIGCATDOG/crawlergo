package main

import (
	"encoding/json"
	"fmt"
)

type hhh struct {
	Name string
}
func main()  {
	tt := hhh{Name: "iii"}
	byte ,_:=json.Marshal(tt)
	fmt.Printf("%s",byte)
	var gg hhh
	json.Unmarshal(byte,&gg)
}
