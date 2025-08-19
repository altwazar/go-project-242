package code

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"path/filepath"
	"testing"
)

// Тест функции GetPathSize с разной комбинацией флагов.
// Не тестируется human из-за маленького размера файлов.
func TestPathSize(t *testing.T) {
	type tc struct {
		name                string
		path                string
		expNonRecNoHidden   string
		expRecNoHidden      string
		expNonRecWithHidden string
		expRecWithHidden    string
	}
	// expNonRecNoHidden - значение без рекурсии и обходом скрытых файлов
	// expRecNoHidden - c рекурсией, но без обхода скрытых файлов
	// expNonRecWithHidden - значение без рекурсии, но с обходом скрытых файлов
	// expRecWithHidden - с рекурсией и обходом скрытых файлов
	cases := []tc{
		{
			name:                "dir_with_only_files",
			path:                filepath.Join("testdata", "dir_with_only_files"),
			expNonRecNoHidden:   "125B",
			expRecNoHidden:      "125B",
			expNonRecWithHidden: "132B",
			expRecWithHidden:    "132B",
		},
		{
			name:                "dir_with_files_and_subdirs",
			path:                filepath.Join("testdata", "dir_with_files_and_subdirs"),
			expNonRecNoHidden:   "125B",
			expRecNoHidden:      "553B",
			expNonRecWithHidden: "132B",
			expRecWithHidden:    "647B",
		},
		{
			name:                "single_file_first_file",
			path:                filepath.Join("testdata", "dir_with_files_and_subdirs", "first_file"),
			expNonRecNoHidden:   "11B",
			expRecNoHidden:      "11B",
			expNonRecWithHidden: "11B",
			expRecWithHidden:    "11B",
		},
		{
			name:                "single_file_large_file",
			path:                filepath.Join("testdata", "dir_with_only_files", "large_file"),
			expNonRecNoHidden:   "114B",
			expRecNoHidden:      "114B",
			expNonRecWithHidden: "114B",
			expRecWithHidden:    "114B",
		},
		{
			name:                "hidden_and_visible",
			path:                filepath.Join("testdata", "hidden_and_visible"),
			expNonRecNoHidden:   "13B",
			expRecNoHidden:      "20B",
			expNonRecWithHidden: "19B",
			expRecWithHidden:    "45B",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			checkPathSize(
				t,
				c.path,
				c.expNonRecNoHidden,
				c.expRecNoHidden,
				c.expNonRecWithHidden,
				c.expRecWithHidden,
				nil,
			)
		})
	}
}

// Отдельно вариант для проверки несуществующего пути
func TestNonexistentPath(t *testing.T) {
	path := filepath.Join("testdata", "___no_such_dir___")
	checkPathSize(
		t,
		path,
		"", "", "", "",
		fs.ErrNotExist,
	)
}

// Проверка функции formatSize
// {передаваемое значение, ожидание в байтах, ожидание в сокращенной форме}
func TestGetPathSizeFile(t *testing.T) {
	cases := []struct {
		size   int64
		expRaw string // human=false
		expHum string // human=true
	}{
		{114, "114B", "114B"},
		{1140, "1140B", "1.1KB"},
		{61140, "61140B", "59.7KB"},
		{50061140, "50061140B", "47.7MB"},
		{30050061140, "30050061140B", "28.0GB"},
		{40030050061140, "40030050061140B", "36.4TB"},
		{70040030050061140, "70040030050061140B", "62.2PB"},
		{9070040030050061140, "9070040030050061140B", "7.9EB"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("size=%d", c.size), func(t *testing.T) {
			gotHum := formatSize(c.size, true)
			gotRaw := formatSize(c.size, false)

			assert.Equal(t, c.expHum, gotHum, "human=true")
			assert.Equal(t, c.expRaw, gotRaw, "human=false")
		})
	}
}

func checkPathSize(
	t *testing.T,
	path string,
	expNonRecNoHidden, expRecNoHidden, expNonRecWithHidden, expRecWithHidden string,
	expErr error,
) {
	t.Helper()

	if expErr != nil {
		modes := []struct {
			rec bool
			all bool
			tag string
		}{
			{false, false, "non-recursive without hidden"},
			{true, false, "recursive without hidden"},
			{false, true, "non-recursive with hidden"},
			{true, true, "recursive with hidden"},
		}
		for _, m := range modes {
			_, err := GetPathSize(path, m.rec, false, m.all)
			if err == nil {
				t.Fatalf("expected error (%s), got nil for %s", expErr, m.tag)
			}
			if !errors.Is(err, expErr) {
				t.Fatalf("expected errors.Is(err, %v) for %s; got err=%v", expErr, m.tag, err)
			}
		}
		return
	}

	// non-recursive, without hidden
	ssize, err := GetPathSize(path, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, expNonRecNoHidden, ssize, "non-recursive without hidden")

	// recursive, without hidden
	ssize, err = GetPathSize(path, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, expRecNoHidden, ssize, "recursive without hidden")

	// non-recursive, with hidden
	ssize, err = GetPathSize(path, false, false, true)
	assert.Nil(t, err)
	assert.Equal(t, expNonRecWithHidden, ssize, "non-recursive with hidden")

	// recursive, with hidden
	ssize, err = GetPathSize(path, true, false, true)
	assert.Nil(t, err)
	assert.Equal(t, expRecWithHidden, ssize, "recursive with hidden")
}
