package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 23
	defaultWidth = 30
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type category string

func (c category) FilterValue() string { return string(c) }

type categoryDelegate struct{}

func (d categoryDelegate) Height() int                             { return 1 }
func (d categoryDelegate) Spacing() int                            { return 0 }
func (d categoryDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d categoryDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(category)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type PrincipalHandler struct {
	list     list.Model
	chosen   string
}

func NewPrincipal() PrincipalHandler {
	categories := []list.Item{
		category("Business"),
		category("Crime"),
		category("Domestic"),
		category("Education"),
		category("Entertainment"),
		category("Environment"),
		category("Food"),
		category("Health"),
		category("Lifestyle"),
		category("Other"),
		category("Politics"),
		category("Science"),
		category("Sports"),
		category("Technology"),
		category("Top"),
		category("Tourism"),
		category("World"),
	}

	l := list.New(categories, categoryDelegate{}, defaultWidth, listHeight)
	l.Title = "Choose a category."
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return PrincipalHandler{
		list: l,
	}
}

func (p PrincipalHandler) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			category := p.list.SelectedItem().FilterValue()
			return p, ChangeStatus(ArticlesTable, NewArticlesTable(category))
		}
	}
	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

func (p PrincipalHandler) Init() tea.Cmd {
	return nil
}

func (p PrincipalHandler) View() string {
	return "\n" + p.list.View()
}
