package main

import (
	"log"
	"os"
	"os/exec"
	"path"

	flag "github.com/spf13/pflag"
)

var (
	ignoreDirs []string
	moduleDirs []string
	basePath   string
	configFile string
)

func init() {
	flag.StringArrayVar(&ignoreDirs, "ignore-dirs", []string{}, "list of ignored directory to traverse into")
	flag.StringVar(&basePath, "base-path", ".", "module base path, default to current directory")
	flag.StringVar(&configFile, "config", ".terraform-docs.yaml", "your terraform-docs yaml config name")
	flag.Parse()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func readDir(basePath string) {
	files, err := os.ReadDir(basePath)
	if err != nil {
		log.Panic(err)
	}
	for _, file := range files {
		if file.IsDir() && !contains(ignoreDirs, file.Name()) {
			newPath := path.Join(basePath, file.Name())
			readDir(newPath)
		} else if file.Name() == "main.tf" {
			moduleDirs = append(moduleDirs, basePath)
		}
	}
}

func main() {

	configPath := path.Join(basePath, configFile)
	if _, err := os.Stat(configPath); err != nil {
		log.Printf("Warning: cannot locate config file %s, proceeding without config file.", configFile)
	}

	readDir(basePath)

	for _, module := range moduleDirs {
		cmd := exec.Command("terraform-docs", "-c", configPath, "markdown", module)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
