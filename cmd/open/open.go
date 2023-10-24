package open

import (
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

		return core.OpenFile(conf, name)
	}

	return c
}
