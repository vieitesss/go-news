package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppStatus int

const (
	Principal AppStatus = iota
)

type News struct {
	Handlers map[AppStatus]tea.Model
	Status   AppStatus
}

func NewNews() News {
	h := make(map[AppStatus]tea.Model)
	h[Principal] = NewPrincipal()

	return News{
		Handlers: h,
		Status:   Principal,
	}
}

func (n News) Init() tea.Cmd {
	return nil
}

func (n News) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	currentStatus := n.Status
	// Global key mappings.
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return n, tea.Quit
		}
	}

	// Per handler key mappings.
	updated, cmd := n.Handlers[n.Status].Update(msg)
	n.Handlers[currentStatus] = updated
	return n, cmd
}

func (n News) View() string {
	return n.Handlers[n.Status].View()
}
