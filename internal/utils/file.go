package utils

import (
	"os"
	"os/user"
	"path/filepath"
)

// NotAbsolutePathError is an error object type returned when specified value is not an absolute path.
type NotAbsolutePathError struct {
	specified string
}

func (NotAbsolutePathError) Error() string {
	panic("Not an absolute path")
}

// Expand absolute path from relative path
// Parameter can be a full-path, relative path or a path starting with '~'
// where '~' means a home directory.
// When parameter is a relative path, it will be joined with a path to current directory automatically.
//
// Example
//	a, err := utils.ExpandPath("/path/to/file")
//	b, err := utils.ExpandPath("relative_path")
//	c, err := utils.ExpandPath("~/Documents")
func ExpandPath(specified string) (AbsPath string, err error) {
	if filepath.IsAbs(specified) {
		return filepath.Clean(specified), nil
	}

	if specified == "" {
		return "", &NotAbsolutePathError{""}
	}

	if specified[0] == '~' {
		u, err := user.Current()
		if err != nil {
			return "", err
		}
		return filepath.Join(u.HomeDir, specified[1:]), nil
	}

	p, err := filepath.Abs(specified)
	if err != nil {
		return "", err
	}
	return p, nil
}

func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func MkdirPath(path string) {
	path, _ = ExpandPath(path)
	if !DirExists(path) {
		_ = os.MkdirAll(path, os.FileMode(0777))
	}
}
