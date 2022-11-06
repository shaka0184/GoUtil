// Package file
// file ファイル関係の処理
package file

import (
	"context"
	"github.com/pkg/errors"
	"github.com/shaka0184/GoUtil/pkg/google/storage"
	"os"
	"path/filepath"
)

func GetFile(filePath string) (string, error) {
	var err error
	res := filePath

	if !Exists(filePath) {
		res, err = os.Executable()
	}

	return res, err
}

// Exists returns int.
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func GetCurrentFileName(fn string) (string, error) {
	if Exists(fn) {
		return fn, nil
	}

	exe, err := os.Executable()
	if err != nil {
		return "", errors.WithStack(err)
	}

	return filepath.Join(filepath.Dir(exe), fn), nil
}

func ReadFile(fn string) ([]byte, error) {
	fName, err := GetCurrentFileName(fn)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(fName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return b, nil
}

func GetLocalOrGcsFile(ctx context.Context, fn, bn string) ([]byte, error) {
	res, err := ReadFile(fn)
	if err != nil {
		return nil, err
	}

	if res != nil && len(res) > 0 {
		return res, nil
	}

	return storage.GetByteSlice(ctx, bn, fn)
}
