package TUI

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"
)

type model struct {
	table table.Model
}

func initialModel() model {
	columns := []table.Column{
		{Title: "Title", Width: 20},
		{Title: "Category", Width: 10},
		{Title: "Rating", Width: 8},
		{Title: "Started", Width: 12},
		{Title: "Finished", Width: 12},
		{Title: "Status", Width: 12},
		{Title: "Notes", Width: 20},
	}

	rows := []table.Row{
		{"Alchemist", "Anime", "100/99", "Sometime", "Sometime", "Finished", "Was good"},
		{"Another Entry", "Manga", "90/100", "Yesterday", "Today", "In Progress", "Interesting!"},
		{"Breaking Bad", "TV Show", "98/100", "Last Year", "Last Week", "Finished", "Amazing plot!"},
		{"Naruto", "Anime", "85/100", "2010", "2015", "Finished", "Classic, but long."},
		{"One Piece", "Anime", "99/100", "2000", "Ongoing", "Ongoing", "The adventure never ends!"},
		{"Death Note", "Anime", "95/100", "2015", "2015", "Finished", "Thrilling!"},
		{"Fullmetal", "Anime", "100/100", "Sometime", "Sometime", "Finished", "Perfect!"},
		{"Attack on Titan", "Anime", "100/100", "2017", "2023", "Finished", "Unbelievable twists!"},
		{"Your Name", "Movie", "92/100", "2018", "2018", "Finished", "Beautiful and emotional."},
		{"Demon Slayer", "Anime", "95/100", "2019", "Ongoing", "Ongoing", "Incredible animation."},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	// Define border styles for the table
	styles := table.DefaultStyles()

	// Header styles
	styles.Header = styles.Header.
		Border(lipgloss.Border{
			Top:         "─",
			Bottom:      "─",
			Left:        "│",
			Right:       "│",
			TopLeft:     "┌",
			TopRight:    "┐",
			BottomLeft:  "└",
			BottomRight: "┘",
		}, true, true, true, true).
		BorderForeground(lipgloss.Color("63")).
		Foreground(lipgloss.Color("231")).
		Bold(true)

	// Cell styles
	styles.Cell = styles.Cell.
		Border(lipgloss.Border{
			Left:  "│",
			Right: "│",
		}, false, true, false, true).
		BorderForeground(lipgloss.Color("63"))

	t.SetStyles(styles)

	return model{table: t}
}

func Tui() {
	zone.NewGlobal()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// Update handles key messages and table navigation
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View renders the table with an outer border
func (m model) View() string {
	// Add an outer border around the entire table
	tableView := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Render(m.table.View())

	return tableView + "\nPress q to quit."
}
