package funcs

import (
	"os"
	"path/filepath"
)

func Get_file_content(working_dir string, file_path string) (string, error) {
	// this function should return the content of the file at the given filepath
	// I should first make an absolute path from the working_dir and then join it with the filepath and make absolute path out of that and check that whether the path exists or not
	absPath, err := filepath.Abs(working_dir)
	if err != nil {
		return "", err
	}
	targetPath, err := filepath.Abs(filepath.Join(absPath, file_path))
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(targetPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
