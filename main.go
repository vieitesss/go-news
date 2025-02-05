package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	tea "github.com/charmbracelet/bubbletea"
)

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Could not read .env file", "error", err)
	}
}

func main() {
	program := tea.NewProgram(
		NewNews(),
		tea.WithAltScreen(),
	)

	if _, err := program.Run(); err != nil {
		os.Exit(1)
	}
}
