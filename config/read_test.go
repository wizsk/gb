package config

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestReadconf(t *testing.T) {
	os.Setenv(rootDirEnvName, "./tmp")
	r := os.Getenv(rootDirEnvName)
	os.Mkdir(r, os.ModePerm)

	df := DefaultConf()
	d, err := yaml.Marshal(df)
	if err != nil {
		t.Fatal(err)
	}

	os.WriteFile(filepath.Join(r, configFileName), d, 0666)
	c, err := readconf()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", c)
}
