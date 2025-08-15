package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize_File(t *testing.T) {
	// testdata/dir_with_only_files/
	size, err := GetSize("testdata/dir_with_only_files/")
	assert.Equal(t, size, int64(132), "they should be equal")
	assert.Nil(t, err)

	// testdata/dir_with_only_subdirs/
	// size, err = GetSize("testdata/dir_with_only_subdirs/")
	// assert.Equal(t, size, int64(0), "they should be equal")
	// assert.Nil(t, err)

	// testdata/dir_with_files_and_subdirs/
	size, err = GetSize("testdata/dir_with_files_and_subdirs/")
	assert.Equal(t, size, int64(132), "they should be equal")
	assert.Nil(t, err)
	// testdata/dir_with_files_and_subdirs/first_file
	size, err = GetSize("testdata/dir_with_files_and_subdirs/first_file")
	assert.Equal(t, size, int64(11), "they should be equal")
	assert.Nil(t, err)

	// testdata/dir_with_only_files/large_file
	size, err = GetSize("testdata/dir_with_only_files/large_file")
	assert.Equal(t, size, int64(114), "they should be equal")
	assert.Nil(t, err)
}
