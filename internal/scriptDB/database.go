package scriptDB

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"strings"

	xdg "github.com/twpayne/go-xdg/v6"
)

func FindDatabase() (*Database, error) {
	dirs, err := xdg.NewBaseDirectorySpecification()
	if err != nil {
		return nil, err
	}

	data_dir := path.Join(dirs.DataHome, "script_manager")
	database := Database{data_dir}

	return &database, nil
}

type Database struct {
	root string
}

func (self *Database) LocalRoot() string {
	return path.Join(self.root, "local")
}

func (self *Database) OnlineRoot() string {
	return path.Join(self.root, "online")
}

func (self *Database) createOnlineScript(service string, user string, name string, contents []byte) error {
	p := path.Join(self.root, "online", service, user, name)
	slog.Info(p)

	err := os.MkdirAll(p, 0o755)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	os.WriteFile(p, contents, 0o755)

	return nil
}

func (self *Database) createLocalScript(name string, contents []byte) error {
	p := path.Join(self.root, "local", name)
	slog.Info(p)

	err := os.MkdirAll(p, 0o755)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	os.WriteFile(p, contents, 0o755)

	return nil
}

func (self *Database) searchOnlineScripts(searchTerm string) (string, error) {
	return "", errors.New("TODO")
}

func (self *Database) list() ([]ScriptListItem, error) {
	local, err := self.ListLocalScripts()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	online, err := self.ListOnlineScripts()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return concat(local, online), nil
}

type ScriptListItem struct {
	Name   string
	OsPath string
}

func (self *ScriptListItem) OnlinePath() string {
	return strings.Join(removeBefore(strings.Split(self.OsPath, "/"), "script_manager"), "")
}

func (self ScriptListItem) String() string {
	return fmt.Sprintf("{%s,%s,%s}", self.Name, self.OnlinePath(), self.OsPath)
}

func (self *Database) RunScript(name string, args []string) error {
	list1, err := self.ListScripts()
	if err != nil {
		return err
	}

	list := arrayfilter(list1, func(item ScriptListItem) bool {
		if item.Name == name {
			return true
		}

		return false
	})
	if len(list) < 1 {
		return errors.New(fmt.Sprintf("Script %s does not exist", name))
	}

	_, err = os.Stat(list[0].OsPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Script %s does not exist", name))
	}

	// AI-Generated
	cmd := exec.Command(list[0].OsPath)

	// Connect the script's stdout to the program's stdout
	cmd.Stdout = os.Stdout

	// Also connect stderr
	cmd.Stderr = os.Stderr

	// Run the command
	return cmd.Run()
}

// Currently can only rename scripts in local
func (self *Database) RenameScript(from string, to string) error {
	list, err := self.ListScripts()
	if err != nil {
		return err
	}

	outputlist := make([]ScriptListItem, 0)
	for _, item := range list {
		if item.Name == from && item.OnlinePath() == "" {
			outputlist = append(outputlist, item)
		}
	}

	if len(outputlist) < 1 {
		return errors.New(fmt.Sprintf("Script \"%s\" does not exist", from))
	}

	parent, _ := path.Split(outputlist[0].OsPath)
	os.Rename(outputlist[0].OsPath, path.Join(parent, to))

	return nil
}

func (self *Database) RemoveLocalScript(name string) error {
	p := path.Join(self.LocalRoot(), name)

	err := os.Remove(p)
	return err
}
