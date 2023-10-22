package newNote

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wizsk/gb/config"
	"github.com/wizsk/gb/core"
)

func Create(c *config.Config, tmpfile *string) *cobra.Command {
	n := &cobra.Command{
		Use:   "new",
		Short: "create a new note",
	}

	var fileName string
	n.Flags().StringVarP(&fileName, "name", "n", "", "name of the new note")

	n.RunE = func(cmd *cobra.Command, args []string) error {
		if fileName == "" && len(args) == 1 {
			fileName = args[0]
		}
		fmt.Println(fileName)

		return core.CreateNewFile(c, fileName, tmpfile)
	}

	return n
}
