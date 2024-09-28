package main

import (
	"Ani-Track/Services"
	"fmt"
)

func main() {
	fmt.Println("Ani-Track")
	content, err := Services.Read()
	if err != nil {
		fmt.Println(err)
	}
	Services.Download("Test", "~/Downloads/file.txt")
	fmt.Println(content)
}
