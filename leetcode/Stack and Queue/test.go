package main

import (
	"fmt"
	"strconv"
)

func main(){
	test := []string{"sd","123"}

	for _,token := range test {
		val ,err := strconv.Atoi(token)
		fmt.Println("%T ,%T\n",val,err)
	}
}
