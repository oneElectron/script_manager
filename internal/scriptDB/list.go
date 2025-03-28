package scriptDB

import (
	"log/slog"
	"os"
	"path"
)

func (self *Database) ListScripts() ([]ScriptListItem, error) {
	list, err := self.list()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (self *Database) ListLocalScripts() ([]ScriptListItem, error) {
	root := self.LocalRoot()
	entries, err := os.ReadDir(root)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	paths := dirEntriesToList(entries, root)

	return paths, nil
}

func (self *Database) ListOnlineScripts() ([]ScriptListItem, error) {
	root := self.OnlineRoot()
	services, err := getSubfolders(root)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	users, err := getAllSubfolders(services)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	scripts, err := getAllSubChildrenEntries(users)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return scripts, nil
}

func getSubfolders(p string) ([]string, error) {
	entries, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}

	folders := arrayfilter(entries, func(e os.DirEntry) bool {
		return e.IsDir()
	})

	paths := arraymap(folders, func(e os.DirEntry) string {
		return path.Join(p, e.Name())
	})

	return paths, nil
}

func getSubchildren(p string) ([]string, error) {
	entries, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}

	paths := arraymap(entries, func(e os.DirEntry) string {
		return path.Join(p, e.Name())
	})

	return paths, nil
}

func getAllSubfolders(paths []string) ([]string, error) {
	output := make([]string, 0)
	for _, p := range paths {
		opaths, err := getSubfolders(p)
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}

		output = concat(output, opaths)
	}

	return output, nil
}

func getAllSubChildren(paths []string) ([]string, error) {
	output := make([]string, 0)
	for _, p := range paths {
		opaths, err := getSubfolders(p)
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}

		output = concat(output, opaths)
	}

	return output, nil
}

func getAllSubChildrenEntries(paths []string) ([]ScriptListItem, error) {
	output := make([]ScriptListItem, 0)
	for _, p := range paths {
		entries, err := os.ReadDir(p)
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}

		rootedEntries := arraymap(entries, func(e os.DirEntry) ScriptListItem {
			return ScriptListItem{
				Name:   e.Name(),
				OsPath: path.Join(p, e.Name()),
			}
		})

		output = concat(output, rootedEntries)
	}

	return output, nil
}
