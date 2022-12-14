package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

type Args struct {
	DirectoryPath string
	Port          string
	FilePath      string
	IsFile        bool
	Theme         string
}

func ParseOSArgs(argv []string) (Args, error) {
	var a Args

	for idx, val := range argv {
		var i string
		if idx != len(argv)-1 {
			i = argv[idx+1]
		}

		switch val {
		case "-p", "--port":
			if !utils.IsValidPort(i) {
				return a, errors.New(fmt.Sprintf("ERROR: %s is not a valid port\n", i))
			}

			a.Port = i

		case "-d", "--dir":
			if !utils.IsValidDirectory(i) {
				return a, errors.New(fmt.Sprintf("ERROR: %s is not valid directory\n", i))
			}

			fullPath, _ := filepath.Abs(i)
			a.DirectoryPath = fullPath

		case "-f", "--file":
			if !utils.IsValidFile(i) {
				return a, errors.New(fmt.Sprintf("ERROR: %s is not a valid file\n", i))
			}

			fullPath, _ := filepath.Abs(i)
			a.FilePath = fullPath
			a.IsFile = true
		case "-t", "--theme":
			if i != "dark" {
				if i != "light" {
					return a, errors.New(fmt.Sprintf("ERROR: %s is not a valid theme\n", i))
				}
			}
			if i != "light" {
				if i != "dark" {
					return a, errors.New(fmt.Sprintf("ERROR: %s is not a valid theme\n", i))
				}
			}
			a.Theme = i
		case "-h", "--help":
			help()
		}
	}

	return a, nil
}

func (a *Args) Parse() Args {
	directoryOnly := a.DirectoryPath != "" && a.FilePath == ""
	fileOnly := a.DirectoryPath == "" && a.FilePath != ""
	dirAndFile := a.DirectoryPath != "" && a.FilePath != ""
	noFileAndNoDirectory := a.DirectoryPath == "" && a.FilePath == ""
	noPort := a.Port == ""
	noArgs := a.Port == "" && a.DirectoryPath == "" && a.FilePath == "" && a.Theme == ""
	noTheme := a.Theme == ""

	if dirAndFile {
		fmt.Println("Found directory and file in arguments only one is allowed. Try -h or --help")
		os.Exit(1)
	}

	if noFileAndNoDirectory {
		p, err := os.Getwd()

		if err != nil {
			fmt.Println("Error resolving cwd. No file or dierctory supplied. Try -h or --help")
			os.Exit(1)
		}
		a.DirectoryPath = p
		a.IsFile = false
		noArgs = true
	}

	if fileOnly {
		a.IsFile = true
	}

	if directoryOnly {
		a.IsFile = false
	}

	if noPort {
		a.Port = "8080"
	}

	if noArgs {
		fmt.Println("No args supplied. Try -h or --help. Continuing with defaults.")
	}

	if noTheme {
		a.Theme = "light"
	}

	return *a
}
