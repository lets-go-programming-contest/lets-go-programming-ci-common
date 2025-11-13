package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	outputTmplFilename  = "task5_test.go"
	undefinedModuleName = "github.com/lgp/not-found"
)

//go:embed task_5.tmpl
var embeddedTemplateData string

var tmpl = template.Must(template.New(string(embeddedTemplateData)).Parse(embeddedTemplateData))

func getModuleNameFromFile(gomodPath string) string {
	modFile, err := os.Open(gomodPath)
	if err != nil {
		panic(fmt.Errorf("open go.mod file: %w", err))
	}
	defer modFile.Close()

	scanner := bufio.NewScanner(modFile)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module"))
		}
	}

	return undefinedModuleName
}

func main() {

	file, err := os.Create(outputTmplFilename)
	if err != nil {
		panic(fmt.Errorf("create output file: %w", err))
	}
	defer file.Close()

	if err := tmpl.Execute(file, getModuleNameFromFile(
		filepath.Join(os.Getenv("SOURCE_DIR"), "go.mod"),
	)); err != nil {
		panic(fmt.Errorf("execute template: %w", err))
	}
}
