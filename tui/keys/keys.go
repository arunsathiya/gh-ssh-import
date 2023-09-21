package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	CursorUp   key.Binding
	CursorDown key.Binding
	Enter      key.Binding
	Cancel     key.Binding
	Quit       key.Binding
	ForceQuit  key.Binding

	State string
}

func (k KeyMap) ShortHelp() []key.Binding {
	var kb []key.Binding

	if k.State != "browsing" {
		kb = append(kb, k.Cancel, k.ForceQuit)
	}

	return kb
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

func NewKeyMap() *KeyMap {
	return &KeyMap{
		CursorUp: key.NewBinding(
			key.WithKeys("ctrl+k"),
			key.WithHelp("ctrl+k", "move up"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("ctrl+j"),
			key.WithHelp("ctrl+j", "move down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Print the SSH key selected"),
		),
		Cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Cancel"),
		),
	}
}
