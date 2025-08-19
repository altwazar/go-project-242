package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize_File(t *testing.T) {

	var ssize string
	var err error

	// testdata/dir_with_only_files/
	// with hidden, not recursive
	ssize, err = GetPathSize("testdata/dir_with_only_files/", false, false, true)
	assert.Equal(t, ssize, "132B", "they should be equal")
	assert.Nil(t, err)

	// without hidden, not recursive
	ssize, err = GetPathSize("testdata/dir_with_only_files/", false, false, false)
	assert.Equal(t, ssize, "125B", "they should be equal")
	assert.Nil(t, err)

	// with hidden, recursive
	ssize, err = GetPathSize("testdata/dir_with_only_files/", true, false, true)
	assert.Equal(t, ssize, "132B", "they should be equal")
	assert.Nil(t, err)

	// without hidden, recursive
	ssize, err = GetPathSize("testdata/dir_with_only_files/", true, false, false)
	assert.Equal(t, ssize, "125B", "they should be equal")
	assert.Nil(t, err)
	// testdata/dir_with_files_and_subdirs/
	// with hidden, not recursiv
	ssize, err = GetPathSize("testdata/dir_with_files_and_subdirs/", false, false, true)
	assert.Equal(t, ssize, "132B", "they should be equal")
	assert.Nil(t, err)

	// without hidden, not recursive
	ssize, err = GetPathSize("testdata/dir_with_files_and_subdirs/", false, false, false)
	assert.Equal(t, ssize, "125B", "they should be equal")
	assert.Nil(t, err)

	// with hidden, recursive
	ssize, err = GetPathSize("testdata/dir_with_files_and_subdirs/", true, false, true)
	assert.Equal(t, ssize, "559B", "they should be equal")
	assert.Nil(t, err)

	// without hidden, recursive
	ssize, err = GetPathSize("testdata/dir_with_files_and_subdirs/", true, false, false)
	assert.Equal(t, ssize, "552B", "they should be equal")
	assert.Nil(t, err)
	// testdata/dir_with_files_and_subdirs/first_file
	// with hidden, not recursive
	ssize, err = GetPathSize("testdata/dir_with_files_and_subdirs/first_file", false, false, true)
	assert.Equal(t, ssize, "11B", "they should be equal")
	assert.Nil(t, err)

	// testdata/dir_with_only_files/large_file
	// with hidden, not recursive
	ssize, err = GetPathSize("testdata/dir_with_only_files/large_file", false, false, true)
	assert.Equal(t, ssize, "114B", "they should be equal")
	assert.Nil(t, err)
	// without hidden, not recursive
	ssize, err = GetPathSize("testdata/dir_with_only_files/large_file", false, false, false)
	assert.Equal(t, ssize, "114B", "they should be equal")
	assert.Nil(t, err)

	// with hidden, recursive
	ssize, err = GetPathSize("testdata/dir_with_only_files/large_file", true, false, true)
	assert.Equal(t, ssize, "114B", "they should be equal")
	assert.Nil(t, err)
	// without hidden, recursive
	ssize, err = GetPathSize("testdata/dir_with_only_files/large_file", true, false, false)
	assert.Equal(t, ssize, "114B", "they should be equal")
	assert.Nil(t, err)

	var size int64
	var fsize string

	// Разные значения размеров
	size = 114
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "114B", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "114B", "they should be equal")

	size = 1140
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "1.1KB", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "1140B", "they should be equal")

	size = 61140
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "59.7KB", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "61140B", "they should be equal")

	size = 50061140
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "47.7MB", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "50061140B", "they should be equal")

	size = 30050061140
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "28.0GB", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "30050061140B", "they should be equal")

	size = 40030050061140
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "36.4TB", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "40030050061140B", "they should be equal")

	size = 40030050061140
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "36.4TB", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "40030050061140B", "they should be equal")

	size = 70040030050061140
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "62.2PB", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "70040030050061140B", "they should be equal")

	size = 9070040030050061140
	fsize = formatSize(size, true)
	assert.Equal(t, fsize, "7.9EB", "they should be equal")

	fsize = formatSize(size, false)
	assert.Equal(t, fsize, "9070040030050061140B", "they should be equal")
}
