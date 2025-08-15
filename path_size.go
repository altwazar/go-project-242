package code

import (
	"os"
)

func GetSize(path string) (int64, error) {
	pinfo, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}
	var size int64 = 0
	// Если директория, то перебор файлов в ней
	if pinfo.IsDir() {
		files, err := os.ReadDir(path)

		if err != nil {
			return 0, err
		}
		for _, file := range files {
			finfo, _ := file.Info()
			// Только размер файлов
			if !finfo.IsDir() {
				size = size + finfo.Size()
			}
		}
	} else {
		size = size + pinfo.Size()
	}
	return size, nil
}
