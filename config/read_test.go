package config

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestReadconf(t *testing.T) {
	os.Setenv(RootDirEnvName, "./tmp")
	r := os.Getenv(RootDirEnvName)
	_ = os.Mkdir(r, os.ModePerm)

	df := DefaultConf()
	d, err := yaml.Marshal(df)
	if err != nil {
		t.Fatal(err)
	}

	_ = os.WriteFile(filepath.Join(r, ConfigFileName), d, 0666)
	c, err := Readconf()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", c)
}
