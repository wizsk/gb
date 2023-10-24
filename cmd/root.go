package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wizsk/gb/cmd/initGb"
	"github.com/wizsk/gb/cmd/newNote"
	"github.com/wizsk/gb/cmd/open"
)

func RootCmd(tmpfile *string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "gb",
		Short: "gb is notetaking cli app",
		Long: `it does

lot of sufffff`,

		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(initGb.InitGb(), open.Open(), newNote.Create())

	return cmd, nil
}
