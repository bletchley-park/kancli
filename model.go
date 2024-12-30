package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	focused status
	lists   []list.Model
	err     error
	loaded  bool
}

func New() *Model {
	return &Model{}
}
func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "Do something", description: "Do something pls"},
		Task{status: todo, title: "Do something2", description: "Do something pls2"},
		Task{status: todo, title: "Do something3", description: "Do something pls3"},
	})

	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: todo, title: "Doing something", description: "Doing something pls"},
		Task{status: todo, title: "Doing something2", description: "Doing something pls2"},
		Task{status: todo, title: "Doing something3", description: "Doing something pls3"},
	})

	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: todo, title: "Did something", description: "Did something pls"},
		Task{status: todo, title: "Did something2", description: "Did something pls2"},
		Task{status: todo, title: "Did something3", description: "Did something pls3"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if !m.loaded {
		return "loading..."
	}
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.lists[todo].View(),
		m.lists[inProgress].View(),
		m.lists[done].View(),
	)
}