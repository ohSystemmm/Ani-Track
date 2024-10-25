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
	table  table.Model
	width  int
	height int
}

func initialModel() model {
	columns := []table.Column{
		{Title: "Title", Width: 20},
		{Title: "Category", Width: 8},
		{Title: "Rating", Width: 8},
		{Title: "Started", Width: 14},
		{Title: "Finished", Width: 14},
		{Title: "Status", Width: 12},
		{Title: "Notes", Width: 40},
	}

	rows := []table.Row{
		{"1 2", "Anime", "100/99", "Sometime", "Sometime", "Finished", "Was good"},
		{"1 2 Entry", "Manga", "90/100", "Yesterday", "Today", "In Progress", "Interesting!"},
		{"1 2 Bad", "TV Show", "98/100", "Last Year", "Last Week", "Finished", "Amazing plot!"},
		{"1 2", "Anime", "85/100", "2010", "2015", "Finished", "Classic, but long."},
		{"1 2 Piece", "Anime", "99/100", "2000", "Ongoing", "Ongoing", "The adventure never ends!"},
		{"1 2 Note", "Anime", "95/100", "2015", "2015", "Finished", "Thrilling!"},
		{"1 2", "Anime", "100/100", "Sometime", "Sometime", "Finished", "Perfect!"},
		{"1 2 on Titan", "Anime", "100/100", "2017", "2023", "Finished", "Unbelievable twists!"},
		{"1 2 Name", "Movie", "92/100", "2018", "2018", "Finished", "Beautiful and emotional."},
		{"1 2 Slayer", "Anime", "95/100", "2019", "Ongoing", "Ongoing", "Incredible animation."},
		{"1 2", "Anime", "100/99", "Sometime", "Sometime", "Finished", "Was good"},
		{"1 2 Entry", "Manga", "90/100", "Yesterday", "Today", "In Progress", "Interesting!"},
		{"1 2 Bad", "TV Show", "98/100", "Last Year", "Last Week", "Finished", "Amazing plot!"},
		{"1 2", "Anime", "85/100", "2010", "2015", "Finished", "Classic, but long."},
		{"1 2 Piece", "Anime", "99/100", "2000", "Ongoing", "Ongoing", "The adventure never ends!"},
		{"1 2 Note", "Anime", "95/100", "2015", "2015", "Finished", "Thrilling!"},
		{"1 2", "Anime", "100/100", "Sometime", "Sometime", "Finished", "Perfect!"},
		{"1 2 on Titan", "Anime", "100/100", "2017", "2023", "Finished", "Unbelievable twists!"},
		{"1 2 Name", "Movie", "92/100", "2018", "2018", "Finished", "Beautiful and emotional."},
		{"1 2 Slayer", "Anime", "95/100", "2019", "Ongoing", "Ongoing", "Incredible animation."},
		{"1 2", "Anime", "100/99", "Sometime", "Sometime", "Finished", "Was good"},
		{"1 2 Entry", "Manga", "90/100", "Yesterday", "Today", "In Progress", "Interesting!"},
		{"1 2 Bad", "TV Show", "98/100", "Last Year", "Last Week", "Finished", "Amazing plot!"},
		{"1 2", "Anime", "85/100", "2010", "2015", "Finished", "Classic, but long."},
		{"1 2 4 Piece", "Anime", "99/100", "2000", "Ongoing", "Ongoing", "The adventure never ends!"},
		{"1 2 4 Note", "Anime", "95/100", "2015", "2015", "Finished", "Thrilling!"},
		{"1 2 4", "Anime", "100/100", "Sometime", "Sometime", "Finished", "Perfect!"},
		{"1 2 4 on Titan", "Anime", "100/100", "2017", "2023", "Finished", "Unbelievable twists!"},
		{"1 2 4 Name", "Movie", "92/100", "2018", "2018", "Finished", "Beautiful and emotional."},
		{"1 2 4 Slayer", "Anime", "95/100", "2019", "Ongoing", "Ongoing", "Incredible animation."},
		{"1 2 4", "Anime", "100/99", "Sometime", "Sometime", "Finished", "Was good"},
		{"1 2 4 Entry", "Manga", "90/100", "Yesterday", "Today", "In Progress", "Interesting!"},
		{"1 2 4 Bad", "TV Show", "98/100", "Last Year", "Last Week", "Finished", "Amazing plot!"},
		{"1 2 4", "Anime", "85/100", "2010", "2015", "Finished", "Classic, but long."},
		{"1 2 4 Piece", "Anime", "99/100", "2000", "Ongoing", "Ongoing", "The adventure never ends!"},
		{"1 2 Note", "Anime", "95/100", "2015", "2015", "Finished", "Thrilling!"},
		{"1 2", "Anime", "100/100", "Sometime", "Sometime", "Finished", "Perfect!"},
		{"1 2 on Titan", "Anime", "100/100", "2017", "2023", "Finished", "Unbelievable twists!"},
		{"1 2 Name", "Movie", "92/100", "2018", "2018", "Finished", "Beautiful and emotional."},
		{"1 2 3 Slayer", "Anime", "95/100", "2019", "Ongoing", "Ongoing", "Incredible animation."},
		{"1 2 3", "Anime", "100/99", "Sometime", "Sometime", "Finished", "Was good"},
		{"1 2 3 Entry", "Manga", "90/100", "Yesterday", "Today", "In Progress", "Interesting!"},
		{"1 2 3 Bad", "TV Show", "98/100", "Last Year", "Last Week", "Finished", "Amazing plot!"},
		{"1 2 3", "Anime", "85/100", "2010", "2015", "Finished", "Classic, but long."},
		{"1 2 3 Piece", "Anime", "99/100", "2000", "Ongoing", "Ongoing", "The adventure never ends!"},
		{"1 2 3 Note", "Anime", "95/100", "2015", "2015", "Finished", "Thrilling!"},
		{"1 2 3", "Anime", "100/100", "Sometime", "Sometime", "Finished", "Perfect!"},
		{"1 2 3 on Titan", "Anime", "100/100", "2017", "2023", "Finished", "Unbelievable twists!"},
		{"1 2 3 Name", "Movie", "92/100", "2018", "2018", "Finished", "Beautiful and emotional."},
		{"1 2 3 Slayer", "Anime", "95/100", "2019", "Ongoing", "Ongoing", "Incredible animation."},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(32),
	)

	// Define border styles for the table
	styles := table.DefaultStyles()

	// Define the selected row style with a highlighted background
	styles.Selected = styles.Selected.
		Foreground(lipgloss.Color("230")). // Text color
		Background(lipgloss.Color("63")).  // Highlight background (blueish)
		Bold(true)                         // Make selected row bold for emphasis

	styles.Header = styles.Header.Bold(true).Background(lipgloss.Color("60"))

	// Apply the styles to the table
	t.SetStyles(styles)

	return model{table: t}
}

func Tui() {
	bzone.NewGlobal()
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Ani-Track"),
	)
}

// Update handles key messages and table navigation
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up": // Scroll up in the table
			m.table.MoveUp(1)
		case "down": // Scroll down in the table
			m.table.MoveDown(1)
		case "n":

		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table.SetHeight((m.height / 2) - 3)
		// case tea.MouseMsg: // Handle mouse events
		// 	mouseEvent := msg.Mouse()  // Get the underlying mouse event
		// 	switch mouseEvent.Button { // Use a switch statement to check mouse button
		// 	case tea.MouseWheelUp:
		// 		m.table.MoveUp(1)
		// 	case tea.MouseWheelDown:
		// 		m.table.MoveDown(1)
		// case tea.MouseMsg:
		// 	// Check for mouse wheel events
		// 	switch msg.Button {
		// 	case msg.Button.MouseWheelUp:
		// 		m.table.MoveUp(1) // Scroll up when the wheel is moved up
		// 	case msg.Button.MouseWheelDown:
		// 		m.table.MoveDown(1) // Scroll down when the wheel is moved down
		// }
	}
	return m, cmd
}

// View renders the table with an outer border
func (m model) View() string {
	// Add an outer border around the entire table
	tableView := lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Render(m.table.View())

	return tableView
}
