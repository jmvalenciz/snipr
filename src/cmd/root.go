package cmd

import (
	"log"
	"os"
	"snipr/src/controller"
	"snipr/src/repository"
	"snipr/src/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snipr",
	Short: "Manage your own snippets",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		repository, err := repository.NewSnippetRepository("./example/db.sqlite")
		if err != nil {
			return err
		}
		controller := controller.NewSnippetController(repository)
		program := tea.NewProgram(ui.NewModel(&controller))
		_, err = program.Run()
		return err
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(os.Stderr, err)
		os.Exit(1)
	}
}
