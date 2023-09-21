package ssh

import (
	"os"
	"path/filepath"
)

type Key struct {
	Name    string
	Path    string
	Content string
}

func GetLocalSSHKeys() (keys []Key, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	sshDir := filepath.Join(homeDir, ".ssh")
	files, err := os.ReadDir(sshDir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".pub" {
			filepath := filepath.Join(sshDir, file.Name())
			keyBytes, err := os.ReadFile(filepath)
			if err != nil {
				return nil, err
			}
			keys = append(keys, Key{
				Name:    file.Name(),
				Path:    filepath,
				Content: string(keyBytes),
			})
		}
	}

	return keys, nil
}
