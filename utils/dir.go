package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetAppBaseDir get the execute file's absolute path
func GetAppBaseDir() string {
	file, _ := exec.LookPath(os.Args[0])
	fpath, _ := filepath.Abs(file)
	name := filepath.Base(file)
	appDir := strings.TrimRight(fpath, name)
	//fmt.Println("[BaseDir]", appDir)

	// appDir := filepath.Base(os.Args[0])
	return appDir
}

func IsPathExisting(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
