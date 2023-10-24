package core

import (
	"os"
	"os/exec"
)

const (
	readWritePermission = 0666
)

func openEditor(file, editor string) error {
	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
