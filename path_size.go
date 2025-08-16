package code

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(path string, recursive bool, human bool, all bool) (string, error) {
	var total int64
	dirs := []string{path}
	pinfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if !pinfo.IsDir() {
		total += pinfo.Size()
		return FormatSize(total, human), nil
	}

	for len(dirs) > 0 {
		dir := dirs[len(dirs)-1]
		dirs = dirs[:len(dirs)-1]

		entries, err := os.ReadDir(dir)
		if err != nil {
			return FormatSize(total, human), err
		}

		for _, entry := range entries {
			if strings.HasPrefix(entry.Name(), ".") && !all {
				continue
			}
			fullPath := dir + "/" + entry.Name()

			info, err := os.Stat(fullPath)
			if err != nil {
				return FormatSize(total, human), err
			}
			if !entry.IsDir() {
				total += info.Size()
			} else if entry.IsDir() && recursive {
				dirs = append(dirs, fullPath)
			}
		}
	}
	return FormatSize(total, human), nil
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
