package main

import (
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

type AppStatus int

const (
	Principal AppStatus = iota
	ArticlesTable
)

type MsgStatusChanged struct {
	Status AppStatus
	Object tea.Model
}

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

func ChangeStatus(status AppStatus, object tea.Model) tea.Cmd {
	return func() tea.Msg {
		switch status {
		case ArticlesTable:
			switch object := object.(type) {
			case ArticlesTableHandler:
			default:
				panic(fmt.Sprintf("Changing to status %v with args type %v", status, reflect.TypeOf(object)))
			}
		}

		return MsgStatusChanged{
			Status: status,
			Object: object,
		}
	}
}

func (n News) Init() tea.Cmd {
	return n.Handlers[n.Status].Init()
}

func (n News) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Global key mappings.
	switch msg := msg.(type) {
	case MsgStatusChanged:
		newStatus := AppStatus(msg.Status)
		n.Status = newStatus
		if msg.Object != nil {
			n.Handlers[newStatus] = msg.Object
		}

		return n, n.Handlers[newStatus].Init()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return n, tea.Quit
		}
	}

	// Per handler key mappings.
	var cmd tea.Cmd
	n.Handlers[n.Status], cmd = n.Handlers[n.Status].Update(msg)
	return n, cmd
}

func (n News) View() string {
	return n.Handlers[n.Status].View()
}
