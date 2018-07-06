package file_directory_new

import (
	"fmt"
	"path/filepath"
	"os"
	"io/ioutil"
)

func init() {
	fmt.Println("===== List ALl the Files under a Directory =====")

	readDir("/home/comp1/copyTest/source")

}

func readDir(root string) {

	var (
		files []string
		err   error
	)

	// filepath.Walk
	files, err = FilePathWalkDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}


	// ioutil.ReadDir
	files, err = IOReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	//os.File.Readdir
	files, err = OSReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}


func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
