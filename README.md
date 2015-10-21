# Go-Init


The go-init Project provide a tool to create a new go project.

## Dependency
-  [client](https://github.com/codegangsta/cli) 
-  [viper](https://github.com/spf13/viper)
-  [logrus](https://github.com/Sirupsen/logrus)

## How To Install
Get the code
```bash
 git clone https://github.com/evonck/go-init.git
 cd go-init
 godep go install
 cd $GOPATH/bin
 ````
 
 ## How To Use
 
 Usage
 ```bash
 ./go-init directory projectName
 ```
 Example:
 ```bash
./go-init ../ init
```

## Set Up
A configuration fie is available. The ocnfiguration file is used to set up the list of file you want to automaticaly generate the file:
```bash
files:
    fileName:
        from: Path to templateFile
        to: directoryPath
```

The go-init project will automatically update:
- $ProjectName string to the appName
- $Directory with the directory path

## Files
By default the go-init project will create:
- main.go
- .gitignore
- Dockerfile
- docker-compose.yml
- makeflile
- Godeps/godeps.json
- VERSION file
- config.yml