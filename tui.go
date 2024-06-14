package main

import (
	"fmt"
	
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)


type item struct {
	title, desc, url string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
		case " ":
            selectedItem := m.list.SelectedItem()
            if selectedItem != nil {
                i := selectedItem.(item)
                err := OpenURL(i.url)
                if err != nil {
                    fmt.Println("Error opening browser:", err)
                }
            }
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.list.SetSize(msg.Width-h, msg.Height-v)
        m.viewport.Width = msg.Width - h
        m.viewport.Height = msg.Height - v
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)

    // Update viewport content based on the selected item
    selectedItem := m.list.SelectedItem()
    if selectedItem != nil {
        i := selectedItem.(item)
        m.viewport.SetContent(fmt.Sprintf("URL: %s\n\nTitle: %s", i.url, i.title, i.desc))
    } else {
        m.viewport.SetContent("No item selected")
    }

    m.viewport, _ = m.viewport.Update(msg)
    return m, cmd
}


func (m model) View() string {
	terminalWidth := m.list.Width()
	listViewWidth := terminalWidth/2
	m.list.SetWidth(listViewWidth)

	listView := docStyle.Render(m.list.View())
    detailView := docStyle.Render(m.viewport.View())

    return lipgloss.JoinHorizontal(lipgloss.Top, listView, detailView)
}