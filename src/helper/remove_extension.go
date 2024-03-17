package helper

import "path/filepath"

func RemoveExtension(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
