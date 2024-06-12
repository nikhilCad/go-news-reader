package main

import (
	"fmt"
	"os"
	
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mmcdole/gofeed"
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
        // case "enter":
        //     // Update the state to reflect the selected item
        //     selectedItem := m.list.SelectedItem()
        //     if selectedItem != nil {
        //         i := selectedItem.(item)
        //         fmt.Println("Selected URL:", i.url)
        //     }
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.list.SetSize(msg.Width-h, msg.Height-v)
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m model) View() string {
	terminalWidth := m.list.Width()
	listViewWidth := terminalWidth/2
	m.list.SetWidth(listViewWidth)

	listView := docStyle.Render(m.list.View())
    selectedItem := m.list.SelectedItem()
    var detailView string
    if selectedItem != nil {
        i := selectedItem.(item)
        // detailView = docStyle.Render(fmt.Sprintf("URL: %s", i.url))

		detailView = lipgloss.NewStyle().
			Width((terminalWidth-listViewWidth)/1).
			Margin(1, 2).
			Render("URL: ", i.url, "\n", i.title)
			// BorderStyle(lipgloss.RoundedBorder()).
			
    } else {
        detailView = docStyle.Render("No item selected")
    }

    return lipgloss.JoinHorizontal(lipgloss.Top, listView, detailView)
}

func main() {

	urls := []string{
		"http://feeds.twit.tv/twit.xml",
		"https://rss.nytimes.com/services/xml/rss/nyt/World.xml",
	}

	var fp = gofeed.NewParser()

	items := []list.Item{}
	
	for _, url := range urls {

		var feedParsed, _ = fp.ParseURL(url)

		fmt.Println(feedParsed.Title)

		for _, parsedItem := range feedParsed.Items {

			items = append(items, item{
				title: parsedItem.Title,
				desc: parsedItem.Description,
				url:   parsedItem.Link,
			})

		}
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "News"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	
}