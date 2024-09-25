package main

import (
	"Ani-Track/Services"
	"Ani-Track/TUI"
	"fmt"
)

func main() {
	output, _ := Services.Read()
	fmt.Println(output)
	TUI.Tui()
}
