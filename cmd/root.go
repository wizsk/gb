package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wizsk/gb/cmd/newNote"
	"github.com/wizsk/gb/config"
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

	cnf, err := config.DefaultConf()
	if err != nil {
		return nil, err
	}

	cmd.AddCommand(newNote.Create(cnf, tmpfile))

	return cmd, nil
}
