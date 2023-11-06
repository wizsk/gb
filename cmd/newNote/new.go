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
			fileName = time.Now().Format("03-04-05-PM_02-01-06")
		}

		err = core.NewNote(conf, fileName)
		if err != nil {
			return err
		}

		db, err := core.GetDefautNotebookDb(conf)
		if err != nil {
			return err
		}

		db.Notes = append(db.Notes, core.Note{Name: fileName, Created: time.Now(), LastModified: time.Now()})
		return db.WriteJson(conf.RootDir, conf.DefaltNoteBook)
	}

	return n
}
