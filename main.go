package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var dirPath string

	if len(os.Args) < 2 {
		fmt.Println("Enter the directory path:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			dirPath = scanner.Text()
		}
		if dirPath == "" {
			fmt.Println("No directory path provided. Exiting...")
			return
		}
	} else {
		dirPath = os.Args[1]
	}

	fmt.Println("Organizing file in:", dirPath)
	files, err := readFiles(dirPath)
	if err != nil {
		fmt.Printf("Error reading the directory %v\n", err)
		return
	}
	for _, file := range files {
		if !file.IsDir() {
			category := getCategory(file.Name())
			categoryPath := filepath.Join(dirPath, category)
			err := createDir(dirPath, category)
			if err != nil {
				fmt.Println("Not able to create the folder:", err)
				continue
			}
			srcPath := filepath.Join(dirPath, file.Name())
			err = moveFile(srcPath, categoryPath)
			if err != nil {
				fmt.Println("Error moveing the file for X resion: ", err)
			} else {
				fmt.Printf("Moved: %s -> %s \n", file.Name(), category)
			}
		}
	}
}

func readFiles(dirPath string) ([]os.DirEntry, error) {
	file, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Not able to read the file from the given directory")
		return nil, err
	}
	return file, nil
}

func getCategory(fileName string) string {
	ext := filepath.Ext(fileName)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return "Images"
	case ".mov", ".mkv", ".mp4":
		return "Videos"
	case ".txt", ".doc", ".pdf":
		return "Documents"
	default:
		return "Others"
	}
}

func createDir(dirPath, category string) error {
	categoryPath := filepath.Join(dirPath, category)
	return os.MkdirAll(categoryPath, os.ModePerm)
}

func moveFile(filePath, destDir string) error {
	destPath := filepath.Join(destDir, filepath.Base(filePath))
	return os.Rename(filePath, destPath)
}
