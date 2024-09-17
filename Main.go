package main

import (
	"Ani-Track/Services"
	"fmt"
	btea "github.com/charmbracelet/bubbletea"
	// bzone "github.com/lrstanley/bubblezone"
	"os"
)

func main() {
	output, _ := Services.Read()
	fmt.Println(output)
	p := btea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
