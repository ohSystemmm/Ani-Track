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

	err = Services.Download("Test", "~/Downloads/file.txt")
	if err != nil {
		fmt.Println(err)
	}

  fmt.Println(content)
}
