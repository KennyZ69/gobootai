package funcs

import (
	"fmt"
	"os"
	"path/filepath"
)

func Get_files_info(working_dir string, dir string) (string, error) {
	// this function should return a string representing the files in the directory
	// I should first make an absolute path from the working_dir and then join it with the dir and make absolute path out of that and check that whether the path exists or not
	absPath, err := filepath.Abs(working_dir)
	if err != nil {
		return "", err
	}
	targetPath, err := filepath.Abs(filepath.Join(absPath, dir))
	if err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(targetPath)
	if err != nil {
		return "", err
	}
	if !fileInfo.IsDir() {
		return "", fmt.Errorf("the path %s is not a directory", targetPath)
	}

	files, err := os.ReadDir(targetPath)
	if err != nil {
		return "", fmt.Errorf("error reading directory %s: %v\n", targetPath, err)
	}

	var resp string
	for _, file := range files {
		filePath := filepath.Join(targetPath, file.Name())
		info, err := os.Stat(filePath)
		if err != nil {
			return "", fmt.Errorf("error getting info for file %s: %v\n", filePath, err)
		}
		resp += fmt.Sprintf("File: %s, Size: %d bytes, IsDir: %v\n", file.Name(), info.Size(), info.IsDir())
	}

	return resp, nil
}
