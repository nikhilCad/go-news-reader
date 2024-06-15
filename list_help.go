package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/key"
)


type customDelegate struct {
    list.DefaultDelegate
}

func newCustomDelegate() customDelegate {
    return customDelegate{
        DefaultDelegate: list.NewDefaultDelegate(),
    }
}

func (d customDelegate) ShortHelp() []key.Binding {
    return []key.Binding{
        key.NewBinding(
            key.WithKeys("ctrl+c"),
            key.WithHelp("ctrl+c", "quit"),
        ),
        key.NewBinding(
            key.WithKeys("space"),
            key.WithHelp("space", "open URL"),
        ),
        key.NewBinding(
            key.WithKeys("a"),
            key.WithHelp("a", "add feed"),
        ),
    }
}

func (d customDelegate) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {
            key.NewBinding(
                key.WithKeys("ctrl+c"),
                key.WithHelp("ctrl+c", "quit"),
            ),
            key.NewBinding(
                key.WithKeys("space"),
                key.WithHelp("space", "open URL"),
            ),
            key.NewBinding(
                key.WithKeys("a"),
                key.WithHelp("a", "add feed"),
            ),
        },
    }
}