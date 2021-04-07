package main

import (
	"fmt"
	"regexp"
)

func main(){
	input := "a<script>alert('sa')</script>@test.com"
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	fmt.Printf("\n email => %v : %v", input, re.MatchString(input))
}