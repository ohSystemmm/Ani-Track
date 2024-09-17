package main

import (
	"Ani-Track/Services"
	"fmt"
)

func main() {
	output, _ := Services.Read()
	fmt.Println()
	fmt.Println(output)
	fmt.Println()
	/*
		p := btea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	*/
}
