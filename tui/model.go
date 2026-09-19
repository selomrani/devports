package tui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/soufyane/devports/scanner"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table    table.Model
	message  string
	err      error
}

type tickMsg time.Time
type updatePortsMsg []scanner.PortInfo

func InitialModel() model {
	columns := []table.Column{
		{Title: "PORT", Width: 10},
		{Title: "PID", Width: 10},
		{Title: "PROCESS", Width: 20},
		{Title: "MEMORY", Width: 15},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return model{
		table:   t,
		message: "Fetching ports...",
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		fetchPorts(),
	)
}

func fetchPorts() tea.Cmd {
	return func() tea.Msg {
		ports, err := scanner.GetListeningPorts()
		if err != nil {
			return err
		}
		return updatePortsMsg(ports)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.message = "Refreshing..."
			return m, fetchPorts()
		case "k":
			if len(m.table.Rows()) > 0 {
				row := m.table.SelectedRow()
				pidStr := row[1]
				pid, err := strconv.ParseInt(pidStr, 10, 32)
				if err == nil && pid > 0 {
					err := scanner.KillProcess(int32(pid))
					if err != nil {
						m.message = fmt.Sprintf("Failed to kill PID %d: %v", pid, err)
					} else {
						m.message = fmt.Sprintf("Successfully killed PID %d", pid)
						return m, fetchPorts()
					}
				}
			}
		}

	case updatePortsMsg:
		var rows []table.Row
		for _, p := range msg {
			memStr := fmt.Sprintf("%.2f%%", p.Memory)
			if p.Memory == 0 {
				memStr = "-"
			}
			rows = append(rows, table.Row{
				fmt.Sprintf("%d", p.Port),
				fmt.Sprintf("%d", p.PID),
				p.Process,
				memStr,
			})
		}
		m.table.SetRows(rows)
		m.message = fmt.Sprintf("Found %d listening ports", len(rows))

	case error:
		m.err = msg
		m.message = fmt.Sprintf("Error: %v", msg)
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	view := baseStyle.Render(m.table.View())
	helpText := "\n  [k] kill process   [r] refresh   [q] quit"
	
	statusText := "\n  Status: " + m.message
	if m.err != nil {
		statusText = "\n  Error: " + m.err.Error()
	}

	return view + statusText + helpText + "\n"
}
