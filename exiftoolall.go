package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Contains[T comparable](A []T, e T) bool {
	for _, v := range A {
		if v == e {
			return true
		}
	}
	return false
}

func main() {
	// get CWD
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// read files on CWD
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	allowedFiletypes := []string{".jpg", ".jpeg", ".png"}
	args := []string{"-overwrite_original", "-all="}
	argsDefaultLen := len(args)

	// append image files in CWD to arguments list
	for _, file := range files {
		fileExtension := filepath.Ext(file.Name())

		if Contains(allowedFiletypes, strings.ToLower(fileExtension)) {
			filePath := filepath.Join(path, file.Name())
			args = append(args, filePath)
		}
	}

	// don't run ExifTool if no image files in directory
	if len(args) == argsDefaultLen {
		log.Println("No valid files in directory")
		os.Exit(0)
	}

	// run ExifTool on image files
	cmd, err := exec.Command("exiftool", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(cmd[:]))
}
