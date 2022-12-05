package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type ProjectConfig struct {
	projectName string
	projectPath string
	projectRepo string
	projectType bool
}

func setupParseFlags(w io.Writer, args []string) (ProjectConfig, error) {
	projectConfig := ProjectConfig{}
	flagSet := flag.NewFlagSet("flasgset", flag.ContinueOnError)
	flagSet.StringVar(&projectConfig.projectName, "n", "",
		"Project name")
	flagSet.StringVar(&projectConfig.projectPath, "d", "",
		"Project location on disk")
	flagSet.StringVar(&projectConfig.projectRepo, "r", "",
		"Project remote repository URL")
	flagSet.BoolVar(&projectConfig.projectType, "s", false,
		"Project will have static assets or not")
	err := flagSet.Parse(args)
	if err != nil {
		log.Fatalf("Cannot parse user input: %s", err)
		return ProjectConfig{}, err
	}
	if err != nil {
		return ProjectConfig{}, err
	}
	return projectConfig, nil

}

func validateConf(conf ProjectConfig) []error {
	errors := make([]error, 0)
	if conf.projectName == "" {
		projectTypeErr := fmt.Errorf("project name cannot br empty")
		errors = append(errors, projectTypeErr)
	}
	if conf.projectRepo == "" {
		projectReportErr := fmt.Errorf("project repository URL cannot be empty")
		errors = append(errors, projectReportErr)
	}
	if conf.projectPath == "" {
		projectPathErr := fmt.Errorf("project path cannot be empty")
		errors = append(errors, projectPathErr)
	}
	return errors

}

func generateScaffold(w io.Writer, conf ProjectConfig) {
	_, err := fmt.Fprintf(w, "🚀🚀 Generating scaffold for project %s in %s\n", conf.projectName, conf.projectPath)
	if err != nil {
		return
	}
}

func main() {

	projectConfig, err := setupParseFlags(os.Stdout, os.Args[1:])
	if err != nil {
		return
	}
	validationErrors := validateConf(projectConfig)
	for _, validationError := range validationErrors {
		_, err := fmt.Fprintf(os.Stdout, validationError.Error())
		fmt.Fprintf(os.Stdout, "\n")
		if err != nil {
			return
		}
	}
	if len(validationErrors) == 0 {
		generateScaffold(os.Stdout, projectConfig)
	}
}
