package code

import (
	"fmt"
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

// Функция форматирования размера
func FormatSize(size int64, human bool) string {
	suf := "B"
	var out string
	if !human {
		out = fmt.Sprintf("%d%s", size, suf)
		return out
	}
	// Список размерностей
	suflist := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	fsize := float64(size)
	// Вычисление подходящего суффикса и значения для размера
	for i := range suflist {
		if fsize < 1024 {
			suf = suflist[i]
			break
		} else {
			fsize = fsize / 1024
		}
	}
	// Если вывод в байтах, то без без знака после запятой
	if suf == "B" {
		out = fmt.Sprintf("%.0f%s", fsize, suf)
	} else {
		out = fmt.Sprintf("%.1f%s", fsize, suf)
	}
	return out
}
