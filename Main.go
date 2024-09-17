package main

import (
	"Ani-Track/Services"
	"fmt"
)

func main() {
	output := Services.Read()
	fmt.Println(output)
}
