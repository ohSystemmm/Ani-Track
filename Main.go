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

	fmt.Println(content)
	if !Services.Save(content, "/home/ohsystemmm/Downloads/file.txt") {
		fmt.Println("Nah I'd win")
	}
}
