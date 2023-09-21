package tui

import (
	"fmt"
	"io"

	"github.com/arunsathiya/gh-ssh-import/tui/keys"
	"github.com/arunsathiya/gh-ssh-import/tui/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type itemDelegate struct {
	keys   *keys.KeyMap
	styles *styles.Styles
}

func newItemDelegate(keys *keys.KeyMap, styles *styles.Styles) *itemDelegate {
	return &itemDelegate{
		keys:   keys,
		styles: styles,
	}
}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	title := d.styles.NormalTitle.Render
	desc := d.styles.NormalDesc.Render

	if index == m.Index() {
		title = func(s ...string) string {
			return d.styles.SelectedTitle.Render("> ", s[0])
		}
		desc = func(s ...string) string {
			return d.styles.SelectedDesc.Render(s[0])
		}
	}

	name := title(i.Name)
	path := desc(i.Path)

	itemListStyle := fmt.Sprintf("%s %s", name, path)

	fmt.Fprint(w, itemListStyle)
}

func (d itemDelegate) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (d itemDelegate) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{d.keys.CursorUp, d.keys.CursorDown, d.keys.Enter, d.keys.Cancel, d.keys.Quit, d.keys.ForceQuit},
	}
}
