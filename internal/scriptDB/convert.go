package scriptDB

import (
	"errors"
	"log/slog"
	"os"
)

func (self *Database) ConvertLocalToOnline(name string, service string, user string) error {
	script, err := self.FindLocalScript(name)
	if err != nil {
		return err
	}
	if script.OnlinePath() != "local" {
		slog.Info(script.OnlinePath())
		err = errors.New("Could not find script locally")
		slog.Error(err.Error())
		return err
	}

	contents, err := os.ReadFile(script.OsPath)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = self.createOnlineScript(service, user, name, contents)
	if err != nil {
		return err
	}

	err = self.RemoveLocalScript(name)

	return nil
}

func (self *Database) ConvertOnlineToLocal(service string, user string, name string) error {
	script, err := self.GetOnlineScript(service, user, name)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	contents, err := os.ReadFile(script)

	err = self.createLocalScript(name, contents)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	os.Remove(script)

	return nil
}
