package main

import (
	"errors"
	"fmt"
)

func test(arg int)(int, error){
	if arg == 21 {
		return -1, errors.New("its not 21")
	}
	return arg + 1, nil
}

type argError struct{
	arg int
	prob string
}

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func main(){
	for _, i := range []int{7, 24}{
		if r, e := test(i); e != nil {
			fmt.Println("test gg", e)
		}else {
			fmt.Println("test win", r)
		}
	} 
}