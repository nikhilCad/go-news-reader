package main

import (
	"fmt"
	
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/glamour"
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
    textInput  textinput.Model
	showInput  bool
}

func (m model) Init() tea.Cmd {
    // m.textInput = textinput.New()
	// m.textInput.Placeholder = "Type here"
	// m.textInput.Focus()
	// m.textInput.CharLimit = 156
	// m.textInput.Width = 20
    // m.viewport.KeyMap = ViewPortKeyMap()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        case "a":
            m.showInput = !m.showInput
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
        m.viewport.Width = msg.Width/2 - v
        m.viewport.Height = msg.Height - h
        m.viewport.Style = lipgloss.NewStyle().
                            BorderStyle(lipgloss.RoundedBorder()).
                            BorderForeground(lipgloss.Color("62"))
    }

    if m.showInput {
        var cmd tea.Cmd
        m.textInput, cmd = m.textInput.Update(msg)
        return m, cmd
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)

    // Update viewport content based on the selected item
    selectedItem := m.list.SelectedItem()
    if selectedItem != nil {

        i := selectedItem.(item)

        renderer, _ := glamour.NewTermRenderer(
            glamour.WithAutoStyle(),
            glamour.WithWordWrap(m.viewport.Width),
        )
    
        str, _ := renderer.Render(fmt.Sprintf("## URL: %s\n\n# Title: %s \n\n%s", i.url, i.title, ParseUrl(i.url)))

        m.viewport.SetContent(str)
        // TODO: Remove this, scroll focus
        m.viewport.GotoTop()
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

    if m.showInput {
		// return m.textInput.View()
        return lipgloss.JoinHorizontal(lipgloss.Top, listView, m.textInput.View())
	}

    return lipgloss.JoinHorizontal(lipgloss.Top, listView, detailView)
}