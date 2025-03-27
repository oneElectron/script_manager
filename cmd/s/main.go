package main

import (
	"fmt"
	"log/slog"
	"os"
	"slices"

	"github.com/oneElectron/script_manager/internal/scriptDB"
)

func main() {
	db, err := scriptDB.FindDatabase()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	var scriptName = ""
	var argsFirst = 0
	for i, item := range os.Args {
		if i < 1 {
			continue
		}

		if item != "" {
			scriptName = item
			argsFirst = i + 1
		}
	}

	if scriptName == "" {
		fmt.Println("No script name given")
		return
	}

	args := slices.Delete(os.Args, 0, argsFirst)
	err = db.RunScript(scriptName, args)
	if err != nil && err.Error() != "Script does not exist" {
		println(err.Error())
	}
}
