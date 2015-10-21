package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var (
	testFiles []*GeneratedFile
)

func TestGetDirectoryName(t *testing.T) {
	directory = "./"
	appName = "go-init"
	getDirectoryName()
	convey.Convey(" directory should be ./go-init", t, func() {
		convey.So(directory, convey.ShouldEqual, "./go-init")
	})
}

func TestExists(t *testing.T) {
	exist, err := exists("./")
	convey.Convey(" err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey(" exist should be true", t, func() {
		convey.So(exist, convey.ShouldEqual, true)
	})
	exist2, err := exists("./test")
	convey.Convey(" err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey(" exist should be false", t, func() {
		convey.So(exist2, convey.ShouldEqual, false)
	})
}

func TestAddFile(t *testing.T) {
	testFile := &GeneratedFile{
		Name: "main.go",
		From: "./template/mainTemplate.go",
		To:   "./",
	}
	testFiles = append(testFiles, testFile)
	addFile(testFile)
	convey.Convey("files should be equal to testFile", t, func() {
		convey.So(files, convey.ShouldResemble, testFiles)
	})
	files = files[:0]
}

func TestGetFiles(t *testing.T) {
	configPath = "./"
	configName = "configTest"
	getFiles()
	convey.Convey("files should be equal to testFile", t, func() {
		convey.So(files, convey.ShouldResemble, testFiles)
	})
}

func TestCreateDirectory(t *testing.T) {
	createDirectory("./test")
	exist, err := exists("./test")
	convey.Convey(" err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey(" exist should be true", t, func() {
		convey.So(exist, convey.ShouldEqual, true)
	})
}

func TestCreateProjectFile(t *testing.T) {
	createProjectFile("./test/.gitignore", "./template/.gitignoreTemplate")
	exist, err := exists("./test/.gitignore")
	convey.Convey(" err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey(" exist should be true", t, func() {
		convey.So(exist, convey.ShouldEqual, true)
	})
}

func TestCopyFile(t *testing.T) {
	copyFile("./test/Dockerfile", "./template/DockerFileTemplate")
	exist, err := exists("./test/Dockerfile")
	convey.Convey(" err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey(" exist should be true", t, func() {
		convey.So(exist, convey.ShouldEqual, true)
	})
}

func TestReplaceVariable(t *testing.T) {
	replaceVariable("./test/Dockerfile", "./test")
	command := exec.Command("grep", "-rnw", "./test/Dockerfile", "-e", "$ProjectName")
	output, _ := command.Output()
	convey.Convey(" exist should be true", t, func() {
		convey.So(len(output), convey.ShouldEqual, 0)
	})
	os.Remove("./test/Dockerfile")
}

func TestIgnoreBinary(t *testing.T) {
	ignoreBinary("./test/.gitignore")
	command := exec.Command("grep", "-rnw", "./test/.gitignore", "-e", "go-init")
	output, _ := command.Output()
	convey.Convey(" exist should be true", t, func() {
		convey.So(len(output), convey.ShouldNotEqual, 0)
	})
	os.Remove("./test/.gitignore")
	os.Remove("./test")
}

func TestGenerate(t *testing.T) {
	directory = "./test"
	Generate(nil)
	exist, err := exists("./test/go-init/main.go")
	convey.Convey(" err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey(" exist should be true", t, func() {
		convey.So(exist, convey.ShouldEqual, true)
	})
	os.Remove("./test/go-init/main.go")
	os.Remove("./test/go-init")
	os.Remove("./test")
}
