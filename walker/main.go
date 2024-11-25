package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	destPath  string
	startPath string
)

// helper function to copy a file
func copyFile(src, dest string) error {
	// open the source file for reading
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return err
	}

	srcInfo, err := file.Stat()
	if err != nil {
		return err
	}
	return os.Chmod(dest, srcInfo.Mode())
}

func walkAction(path string, d os.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if !d.IsDir() {
		// compute the relative path of the file
		relPath, err := filepath.Rel(startPath, path)
		if err != nil {
			return err
		}

		// extract the parent directory name
		dirName := filepath.Base(filepath.Dir(path))

		// construct the new file name with the directory name as a prefix
		newFileName := dirName + "_" + filepath.Base(path)

		// construct the destination path, preserving hierarchy
		destDir := filepath.Join(destPath, filepath.Dir(relPath))
		newPath := filepath.Join(destDir, newFileName)

		// ensure the destination directory exists
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			return err
		}

		// copy the file
		if err := copyFile(path, newPath); err != nil {
			return err
		}

		fmt.Println("Copied:", path, "->", newPath)
	}

	return nil
}

func main() {
	home, _ := os.UserHomeDir()
	destPath = home + "/storage"
	startPath = home + "/files"

	err := filepath.WalkDir(startPath, walkAction)
	if err != nil {
		log.Fatal(err)
	}
}
