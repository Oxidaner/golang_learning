package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := strconv.FormatBool(true)
	fmt.Printf("type:%T value:%#v\n", s1, s1) // type:string value:"true"
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("type:%T value:%#v\n", s2, s2) // type:string value:"3.1415E+00"
	s3 := strconv.FormatInt(-2, 16)
	fmt.Printf("type:%T value:%#v\n", s3, s3) // type:string value:"-2"
	s4 := strconv.FormatUint(2, 16)
	fmt.Printf("type:%T value:%#v\n", s4, s4) // type:string value:"2"
}
