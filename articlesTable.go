package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vieitesss/go-news/api"
)

type MsgGetByCategoryResponse api.ApiResponse

type ArticlesTableHandler struct {
	Category string
	Response api.ApiResponse
}

func NewArticlesTable(category string) ArticlesTableHandler {
	return ArticlesTableHandler{
		Category: category,
	}
}

func GetArticlesByCategory(category string) tea.Cmd {
	return func() tea.Msg {
		return MsgGetByCategoryResponse(api.GetResults(category))
	}
}

func (a ArticlesTableHandler) Init() tea.Cmd {
	return GetArticlesByCategory(a.Category)
}

func (a ArticlesTableHandler) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case MsgGetByCategoryResponse:
		a.Response = api.ApiResponse(msg)
	}
	return a, nil
}

func (a ArticlesTableHandler) View() string {
	return fmt.Sprintf("%+v", a)
}
