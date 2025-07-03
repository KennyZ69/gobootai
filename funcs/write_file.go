package funcs

import (
	"os"
	"path/filepath"
)

func Write_file(working_dir, file_path, content string) (int, error) {
	// this function should write the content to the specified file path and return the number of bytes written along with possible errors
	absPath, err := filepath.Abs(working_dir)
	if err != nil {
		return 0, err
	}
	targetPath, err := filepath.Abs(filepath.Join(absPath, file_path))
	if err != nil {
		return 0, err
	}

	file, err := os.Create(targetPath)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	b, err := file.WriteString(content)

	return b, err
}
