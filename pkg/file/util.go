// Package file
// file ファイル関係の処理
package file

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
