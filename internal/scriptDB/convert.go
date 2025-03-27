package scriptDB

import(
	"errors"
	"os"
	"log/slog"
)

func (self *Database) ConvertLocalToOnline(name string, service string, user string) error {
	script, err := self.FindScript(name)
	if err != nil {
		return err
	}
	if script.OnlinePath() != "" {
		return errors.New("Count not find script locally")
	}

	contents, err := os.ReadFile(script.OsPath)
	if err != nil {
		return err
	}

	err = self.createOnlineScript(service, user, name, contents)
	if err != nil {
		return err
	}

	return nil
}

func (self *Database) ConvertOnlineToLocal(service string, user string, name string) error {
	script, err := self.GetOnlineScript(service, user, name)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	contents, err := os.ReadFile(script)

	self.createLocalScript(name, contents)

	os.Remove(script)

	return nil
}
