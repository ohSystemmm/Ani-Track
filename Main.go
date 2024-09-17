package main

import (
	"Ani-Track/Services"
	"fmt"
)

func main() {
	output, _ := Services.Read()
	fmt.Println(output)
}
