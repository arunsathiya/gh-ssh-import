package main

import (
	"fmt"
	"os"

	"github.com/arunsathiya/gh-ssh-import/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.NewModel())

	err := p.Start()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
