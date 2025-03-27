package scriptDB

import (
	"errors"
	"log/slog"
	"os"
	"path"
	"slices"
)

func (self *Database) GetOnlineScript(service string, user string, name string) (string, error) {
	p := path.Join(self.root, "online", service, user, name)

	_, err := os.Stat(p)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	return p, nil
}

func (self *Database) GetLocalScript(name string) (string, error) {
	p := path.Join(self.root, "local", name)

	_, err := os.Stat(p)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	return p, nil
}

func (self *Database) FindOnlineScript(name string) (string, error) {
	list, err := self.ListOnlineScripts()
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	for _, item := range list {
		if item.Name == name {
			return item.OsPath, nil
		}
	}

	return "", errors.New("script not found")
}

func (self *Database) FindScript(name string) (*ScriptListItem, error) {
	list, err := self.ListScripts()
	if err != nil {
		return nil, err
	}

	items := make([]ScriptListItem, 0)
	for _, item := range list {
		if item.OnlinePath() == "" {
			items = slices.Insert(items, 0, item)
		} else {
			items = append(items, item)
		}
	}

	if len(items) < 1 {
		return nil, errors.New("Could not find script")
	}

	return &items[0], nil
}

func (self *Database) FindLocalScript(name string) (*ScriptListItem, error) {
	p, err := self.GetLocalScript(name)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &ScriptListItem{
		Name:   name,
		OsPath: p,
	}, nil
}
