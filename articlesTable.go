package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vieitesss/go-news/api"
)

type MsgGetByCategoryResponse api.ApiResponse

type ArticlesTableHandler struct {
	spinner  spinner.Model
	waiting  bool
	Category string
	Response api.ApiResponse
}

func NewArticlesTable(category string) ArticlesTableHandler {
	s := spinner.New()
	s.Spinner = spinner.Meter
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return ArticlesTableHandler{
		spinner:  s,
		waiting:  true,
		Category: category,
	}
}

func GetArticlesByCategory(category string) tea.Cmd {
	return func() tea.Msg {
		return MsgGetByCategoryResponse(api.GetResults(category))
	}
}

func (a ArticlesTableHandler) Init() tea.Cmd {
	return tea.Batch(
		GetArticlesByCategory(a.Category),
		a.spinner.Tick,
	)
}

func (a ArticlesTableHandler) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case MsgGetByCategoryResponse:
		a.waiting = false
		a.Response = api.ApiResponse(msg)

	case spinner.TickMsg:
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		return a, cmd
	}

	return a, nil
}

func (a ArticlesTableHandler) View() string {
	if a.waiting {
		return fmt.Sprintf("\n    %s", a.spinner.View())
	}

	return fmt.Sprintf("%+v", a.Response)
}
