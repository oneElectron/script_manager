package edit

import (
	"io"
	"os"
	"path"
	"strings"

	"github.com/confluentinc/go-editor"
)

const defaultScript = `#! /usr/bin/env sh
# Set this shell to whatever you would like.

# Say hello
echo Hello World!!!
`

func Editor(p string) error {
	var r io.Reader = nil

	file, err := os.Open(p)
	if err == nil {
		r = io.Reader(file)
	} else {
		r = strings.NewReader(defaultScript)
	}

	_, filename := path.Split(p)

	edit := editor.NewEditor()
	buf, filepath, err := edit.LaunchTempFile(filename, r)

	print(p)
	os.WriteFile(p, buf, os.ModePerm)

	if err != nil {
		os.Remove(filepath)
		return err
	}

	return nil
}
