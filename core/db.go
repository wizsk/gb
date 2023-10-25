package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/wizsk/gb/config"
)

const dbFileName = "gb.db.json"

type NotesDb struct {
	Notes []Note
}

type Note struct {
	Name         string
	Created      time.Time
	LastModified time.Time `json:"last_modified"`
}

func (nd *NotesDb) UpdateLastModified(conf *config.Config, note string) error {
	for i := 0; i < len(nd.Notes); i++ {
		if nd.Notes[i].Name == note {
			nd.Notes[i].LastModified = time.Now()
			return nd.WriteJson(conf.RootDir, conf.DefaltNoteBook)
		}
	}

	return fmt.Errorf("could not fild %q in the db", note)
}

// root, notebook name
func (nd *NotesDb) WriteJson(root, nb string) error {
	data, err := json.MarshalIndent(nd, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(root, nb, dbFileName), data, readWritePermission)
}

func GetDefautNotebookDb(conf *config.Config) (*NotesDb, error) {
	return GetNotebookDb(conf, conf.DefaltNoteBook)
}

func GetNotebookDb(conf *config.Config, nName string) (*NotesDb, error) {
	if nName == "" {
		return nil, fmt.Errorf("notebook name is empty")
	}
	path := filepath.Join(conf.RootDir, nName, dbFileName)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var nDb NotesDb
	if err = json.Unmarshal(data, &nDb); err != nil {
		return nil, err
	}
	return &nDb, nil
}
