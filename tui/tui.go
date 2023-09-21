package tui

import (
	"fmt"
	"log"

	"github.com/arunsathiya/gh-ssh-import/github"
	"github.com/arunsathiya/gh-ssh-import/ssh"
	"github.com/arunsathiya/gh-ssh-import/tui/keys"
	"github.com/arunsathiya/gh-ssh-import/tui/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	defaultWidth = 20
	listHeight   = 12
)

type item ssh.Key

func (i item) FilterValue() string { return i.Name }

type state int

const (
	browsing state = iota
)

type Model struct {
	keyMap *keys.KeyMap
	list   list.Model
	styles styles.Styles
	state  state
}

func NewModel() Model {
	localSSHKeys, err := ssh.GetLocalSSHKeys()
	if err != nil {
		log.Fatal("something went wrong: %w", err)
	}

	items := []list.Item{}
	for _, localSSHKey := range localSSHKeys {
		items = append(items, item{
			Name: localSSHKey.Name,
			Path: localSSHKey.Path,
		})
	}

	styles := styles.DefaultStyles()
	keys := keys.NewKeyMap()

	l := list.New(items, newItemDelegate(keys, &styles), defaultWidth, listHeight)
	l.Title = "Your Keys"
	l.SetShowStatusBar(true)
	l.Styles.PaginationStyle = styles.Pagination
	l.Styles.HelpStyle = styles.Help

	return Model{
		keyMap: keys,
		list:   l,
		styles: styles,
		state:  browsing,
	}
}

func (m *Model) updateKeybindins() {
	if m.list.SettingFilter() {
		m.keyMap.Enter.SetEnabled(false)
	}

	switch m.state {
	case browsing:
		m.keyMap.Enter.SetEnabled(true)
		m.keyMap.Cancel.SetEnabled(false)

	default:
		m.keyMap.Enter.SetEnabled(true)
		m.keyMap.Cancel.SetEnabled(false)
	}
}

func listUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.list.KeyMap.AcceptWhileFiltering):
			m.state = browsing
			m.updateKeybindins()

		case key.Matches(msg, m.keyMap.CursorUp):
			m.list.CursorUp()

		case key.Matches(msg, m.keyMap.CursorDown):
			m.list.CursorDown()

		case key.Matches(msg, m.keyMap.Enter):
			if i, ok := m.list.SelectedItem().(item); ok {
				uploadResponse, err := github.UploadSshPublicKey(i.Name, i.Path)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println((m.styles.NormalTitle.Copy().MarginTop(1).Render(i.Name, " - ", uploadResponse.String())))
				return m, tea.Quit
			}

		}
	}

	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.list.SettingFilter() {
		m.keyMap.Enter.SetEnabled(false)
	}

	switch m.state {
	case browsing:
		return listUpdate(msg, m)

	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.state {
	case browsing:
		return "\n" + m.list.View()

	default:
		return ""
	}
}
