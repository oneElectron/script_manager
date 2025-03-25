package script_db

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"

	xdg "github.com/twpayne/go-xdg/v6"
)

type database struct {
	root string
}

func newDatabase(root string) database {
	return database{
		root: root,
	}
}

func (self *database) list() ([]ScriptListItem, error) {
	mapper := func(entry fs.DirEntry, p string) ScriptListItem {
		name := entry.Name()
		newpath := remove(removeBefore(strings.Split(p, "/"), "script_manager"), 1)

		return ScriptListItem{
			Name:   name,
			Path:   strings.Join(newpath, "/"),
			osPath: path.Join(p, name),
		}
	}
	localpath := path.Clean(path.Join(self.root, "local"))
	local, err := os.ReadDir(localpath)
	if err != nil {
		os.Mkdir(localpath, fs.ModeDir|fs.ModePerm)
		local, err = os.ReadDir(localpath)
		if err != nil {
			return nil, err
		}
	}

	onlinepath := path.Join(self.root, "online")
	online, err := os.ReadDir(onlinepath)
	if err != nil {
		os.Mkdir(onlinepath, fs.ModeDir|fs.ModePerm)
		online, err = os.ReadDir(onlinepath)

		if err != nil {
			return nil, err
		}
	}

	return concat(
		arraymap(local, func(entry os.DirEntry) ScriptListItem { return mapper(entry, localpath) }),
		arraymap(online, func(entry os.DirEntry) ScriptListItem { return mapper(entry, onlinepath) }),
	), nil
}

type ScriptListItem struct {
	Name   string
	Path   string
	osPath string
}

func (self ScriptListItem) String() string {
	return fmt.Sprintf("{%s,%s,%s}", self.Name, self.Path, self.osPath)
}

func ListFiles() ([]ScriptListItem, error) {
	dirs, err := xdg.NewBaseDirectorySpecification()
	if err != nil {
		return nil, err
	}

	data_dir := path.Join(dirs.DataHome, "script_manager")
	database := newDatabase(data_dir)
	list, err := database.list()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func RunScript(name string, args []string) error {
	list1, err := ListFiles()
	if err != nil {
		return err
	}

	list := arrayfilter(list1, func(item ScriptListItem) bool {
		if item.Name == name {
			return true
		}

		return false
	})

	// AI-Generated
	cmd := exec.Command(list[0].osPath)

	// Connect the script's stdout to the program's stdout
	cmd.Stdout = os.Stdout

	// Also connect stderr
	cmd.Stderr = os.Stderr

	// Run the command
	return cmd.Run()
}
