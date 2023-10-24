package initGb

import (
	"os"
	"testing"

	"github.com/wizsk/gb/config"
)

func TestInit(t *testing.T) {
	os.RemoveAll("./tmp")
	os.Setenv(config.RootDirEnvName, "./tmp")
	getP := func() (string, error) {
		return "hi there", nil
	}

	if err := initGB(true, getP); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
