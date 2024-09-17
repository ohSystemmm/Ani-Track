package main

import (
	"Ani-Track/Services"
	"fmt"
)

func main() {
	test, err := Services.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
}
