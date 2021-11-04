// Package File
// File ファイル関係の処理
package File

import "os"

func getFile(filePath string) (string, error) {
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
