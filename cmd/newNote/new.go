package newNote

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/wizsk/gb/config"
	"github.com/wizsk/gb/core"
)

func Create() *cobra.Command {
	n := &cobra.Command{
		Use:   "new",
		Short: "create a new note",
	}

	var fileName string
	n.Flags().StringVarP(&fileName, "name", "n", "", "name of the new note")

	n.RunE = func(cmd *cobra.Command, args []string) error {
		conf, err := config.Readconf()
		if err != nil {
			return err
		}
		if fileName == "" && len(args) == 1 {
			fileName = args[0]
		}

		if fileName == "" {
			fileName = time.Now().Format("02-01-06_03-04-05-PM")
		}

		return core.NewNote(conf, fileName)
	}

	return n
}
