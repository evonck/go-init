package main

import (
	"io"
	"os"
	"os/exec"
	"path"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var (
	repository string
	namespace  string
	project    string
	once       sync.Once
	files      []*GeneratedFile
	directory  string
)

// Generate generate the go folder
func Generate(ctx *cli.Context) {
	getDirectoryName()
	exist, err := exists(directory)
	if exist && !force {
		log.Fatal("A directory with these name already exist: ", directory)
		log.Fatal("You can use the -f to force the creation of the folder ")
		return
	}
	if exist && force {
		os.RemoveAll(directory)
	}
	if err != nil {
		log.Fatal("An error occur: ", err)
		return
	}
	getFiles()
	for _, file := range files {
		directoryPath := path.Join(directory, file.To)
		filePath := path.Join(directoryPath, file.Name)
		createDirectory(directoryPath)
		createProjectFile(filePath, file.From)
		replaceVariable(filePath, directoryPath)
		if file.Name == ".gitignore" {
			ignoreBinary(filePath)
		}
	}
}

// getDirectoryName get the directory name we need to create
func getDirectoryName() {
	if directory[len(directory)-1:] != "/" {
		directory += "/"
	}
	directory += appName
}

// exists returns whether or not the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// getFiles get the list of file to create from the conifg file
func getFiles() {
	files := Config().GetStringMap("files")
	for fileName, localisation := range files {
		var from string
		var to string
		for key, value := range localisation.(map[interface{}]interface{}) {
			if key.(string) == "from" {
				from = value.(string)
			}
			if key.(string) == "to" {
				to = value.(string)
			}
		}
		addFile(&GeneratedFile{
			Name: fileName,
			From: from,
			To:   to,
		})
	}
}

// addFile to the gernteFile list
func addFile(file *GeneratedFile) {
	if files == nil {
		once.Do(func() {
			files = make([]*GeneratedFile, 0)
		})
	}
	files = append(files, file)
}

// createDirectory create a new directory
func createDirectory(directory string) {
	if err := os.MkdirAll(directory, 0744); err != nil {
		log.Fatal("an error occur during the creation of the folder")
		return
	}
}

// createProjectFile create a file
func createProjectFile(filePath, sourceFileName string) {
	err := copyFile(filePath, sourceFileName)
	if err != nil {
		log.Debug(err)
		log.Fatalf("an error occur during the creation of the %s file, error:", filePath, err)
	}
}

// copyFile copy a file from a destination to a source
func copyFile(destination, source string) error {
	file, err := os.Open(source)
	defer file.Close()
	if err != nil {
		return err
	}
	destinationFile, err := os.Create(destination)
	defer destinationFile.Close()
	if err != nil {
		return err
	}
	if _, err := io.Copy(destinationFile, file); err != nil {
		destinationFile.Close()
		return err
	}
	return nil
}

// replaceVariable replace the $ProjectName by the app Name in the created file
func replaceVariable(filePath, directoryPath string) {
	command := exec.Command("sed", "-i", "", "s/$ProjectName/"+appName+"/g", filePath)
	command.Run()
	command = exec.Command("sed", "-i", "", "s/$Directory/"+directoryPath+"/g", filePath)
	command.Run()
}

// ignoreBinary Add the binary Name to the .gitignore file
func ignoreBinary(filePath string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return
	}
	_, err = f.WriteString("\n")
	_, err = f.WriteString(appName)
}
