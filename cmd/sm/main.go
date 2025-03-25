package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"slices"

	flags "github.com/jessevdk/go-flags"

	"github.com/oneElectron/script_manager/internal/edit"
	"github.com/oneElectron/script_manager/internal/script_db"
	xdg "github.com/twpayne/go-xdg/v6"
)

func main() {
	options := parseOptions()

	dirs, err := xdg.NewBaseDirectorySpecification()
	if err != nil {
		print(err.Error())
		return
	}

	if options.List {
		list, err := script_db.ListFiles()
		if err != nil {
			print(err)
			return
		}

		for _, item := range list {
			if item.Path == "" {
				fmt.Printf("%s (local)", item.Name)
			} else {
				fmt.Printf("%s (%s)", item.Name, item.Path)
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

	args := slices.Delete(os.Args, 0, argsFirst)
	err = script_db.RunScript(scriptName, args)
	if err != nil {
		println(err.Error())
	}

	// _ := path.Join(xdg.DataHome, "script_manager")
}

// --------------------- AI Generated Snippet Start ---------------------
// opts defines the command line options for the CLI.
// Fields correspond to the options in the help menu.
type opts struct {
	Delete string `short:"d" long:"delete" description:"Delete a script"`
	Edit   string `short:"e" long:"edit" description:"Create/Edit a script"`
	List   bool   `short:"l" long:"list" description:"Show list of all scripts"`
	Help   bool   `short:"h" long:"help" description:"Show help"`
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

// runDelete deletes the specified script.
// AI-Generated
func runDelete(script string) {
	// AI-Generated
	fmt.Printf("Deleting script: %s\n", script)
	// Add your deletion logic here.
}

// runEdit creates or edits the specified script.
// AI-Generated
func runEdit(script string) {
	// AI-Generated
	fmt.Printf("Editing/Creating script: %s\n", script)
	// Add your edit/create logic here.
}

// runList lists all available scripts.
// AI-Generated
func runList() {
	// AI-Generated
	fmt.Println("Listing all scripts...")
	// Example scripts list. Replace with your actual logic.
	fmt.Println("Script1")
	fmt.Println("Script2")
}

// --------------------- AI Generated Snippet End ---------------------
