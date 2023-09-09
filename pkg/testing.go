package pkg

import (
	"os"
	"path/filepath"
)

func GetConfigPath(basePath string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", nil
	}
	confPath, err := filepath.Abs(wd + "/" + basePath)
	if err != nil {
		return "", nil
	}

	return confPath + "/config.yaml", nil
}
