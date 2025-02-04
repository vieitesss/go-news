package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type News struct {

}

func NewNews() News {
	return News{

	}
}

func (n News) Init() tea.Cmd {
	return nil
}

func (n News) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return n, nil
}

func (n News) View() string {
	return "Hello world!"
}
