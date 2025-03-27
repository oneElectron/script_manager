package main

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"slices"

	flags "github.com/jessevdk/go-flags"

	"github.com/oneElectron/script_manager/internal/edit"
	"github.com/oneElectron/script_manager/internal/scriptDB"
	xdg "github.com/twpayne/go-xdg/v6"
)


func main() {
	db, err := scriptDB.FindDatabase();
	if err != nil {
		slog.Error(err.Error())
		return
	}
	options := parseOptions()

	dirs, err := xdg.NewBaseDirectorySpecification()
	if err != nil {
		print(err.Error())
		return
	}

	if options.List {
		list, err := db.ListScripts()
		if err != nil {
			print(err)
			return
		}

		for _, item := range list {
			if item.OnlinePath() == "" {
				fmt.Printf("%s (local)\n", item.Name)
			} else {
				fmt.Printf("%s (%s)\n", item.Name, item.OnlinePath())
			}
		}

		return
	} else if options.Edit != "" {
		os.Mkdir(path.Join(dirs.DataDirs[0], "script_manager"), fs.ModeDir|fs.ModePerm)
		err := os.Mkdir(path.Join(dirs.DataDirs[0], "script_manager", "local"), fs.ModeDir|fs.ModePerm)

		err = edit.Editor(path.Join(dirs.DataDirs[0], "script_manager", "local", options.Edit))
		if err != nil {
			print(err.Error())
		}

		return
	} else if options.Delete != "" {
		os.Remove(path.Join(dirs.DataDirs[0], "script_manager", "local", options.Delete))
		return
	} else if len(options.Rename) != 0 {
		for from, to := range options.Rename {
			err = db.RenameScript(from, to)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
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

// --------------------- AI Generated Snippet Start ---------------------
// opts defines the command line options for the CLI.
// Fields correspond to the options in the help menu.
type opts struct {
	Delete string            `short:"d" long:"delete" description:"Delete a script"`
	Edit   string            `short:"e" long:"edit" description:"Create/Edit a script"`
	List   bool              `short:"l" long:"list" description:"Show list of all scripts"`
	Help   bool              `short:"h" long:"help" description:"Show help"`
	Rename map[string]string `long:"rename" description:"Rename a script (format: --rename oldname:newname)"`
}

func parseOptions() opts { // AI-Generated
	var options opts
	parser := flags.NewParser(&options, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
	if options.Help {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}
	return options
}

// --------------------- AI Generated Snippet End ---------------------
