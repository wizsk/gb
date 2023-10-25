package open

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/wizsk/gb/config"
	"github.com/wizsk/gb/core"
)

func Open() *cobra.Command {
	c := &cobra.Command{
		Use:   "open",
		Short: "open a specified note",
	}

	var name string
	var index int
	c.Flags().StringVarP(&name, "name", "n", "", "name of the note. ex: fo")
	c.Flags().IntVarP(&index, "index", "i", -1, "index of the note")

	c.RunE = func(_ *cobra.Command, args []string) error {
		if name != "" && index != -1 {
			return fmt.Errorf("cannot use both -n and -i flags as the same time")
		}

		var err error
		if name == "" && len(args) == 1 {
			if index, err = strconv.Atoi(args[0]); err != nil {
				index = -1
				name = args[0]
			}
		}

		// TODO: add index support
		if index > 0 {
		}

		conf, err := config.Readconf()
		if err != nil {
			return err
		}

		err = core.OpenFile(conf, name)
		if err != nil {
			return err
		}

		db, err := core.GetDefautNotebookDb(conf)
		if err != nil {
			return err
		}
		return db.UpdateLastModified(conf, name)
	}

	return c
}
