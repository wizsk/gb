package init

import (
	"os"
	"testing"

	"github.com/wizsk/gb/config"
)

func TestInit(t *testing.T) {
	os.Setenv(config.RootDirEnvName, "./tmp")
	getP := func() (string, error) {
		return "hi there", nil
	}

	if err := Init(false, getP); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
