package main

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_setupParseFlags(t *testing.T) {
	// Test setup
	testCases := []struct {
		testCaseName   string
		testInput      []string
		expectedConfig ProjectConfig
	}{{"All args set", []string{"-n", "MyProject", "-d", "/path/to/dir", "-r", "github.com/iplitharas/myproject"},
		ProjectConfig{
			projectName: "MyProject",
			projectPath: "/path/to/dir",
			projectRepo: "github.com/iplitharas/myproject",
			projectType: false,
		}}, {"Project name missing", []string{"-d", "/path/to/dir", "-r", "github.com/iplitharas/myproject"},
		ProjectConfig{
			projectName: "",
			projectPath: "/path/to/dir",
			projectRepo: "github.com/iplitharas/myproject",
			projectType: false,
		}}, {"Project path missing", []string{"-n", "MyProject", "-r", "github.com/iplitharas/myproject"},
		ProjectConfig{
			projectName: "MyProject",
			projectPath: "",
			projectRepo: "github.com/iplitharas/myproject",
			projectType: false,
		}},
		{"Project repo missing", []string{"-n", "MyProject", "-d", "./"},
			ProjectConfig{
				projectName: "MyProject",
				projectPath: "./",
				projectRepo: "",
				projectType: false,
			}},
		{"Project with assets", []string{"-n", "MyProject", "-d", "./", "-r",
			"github.com/iplitharas/myproject", "-s"},
			ProjectConfig{
				projectName: "MyProject",
				projectPath: "./",
				projectRepo: "github.com/iplitharas/myproject",
				projectType: true,
			}},
	}
	// When we call the setupParseFlags
	for _, testCase := range testCases {
		projectConfig, err := setupParseFlags(os.Stdout, testCase.testInput)
		if err != nil {
			log.Fatal(err)

		}
		// Then we expect the right ProjectConfig
		if !reflect.DeepEqual(projectConfig, testCase.expectedConfig) {
			t.Errorf("Test case: %v, Expected %v got %v", testCase.testCaseName, testCase.expectedConfig, projectConfig)
		}

	}

}

func Test_validateConf(t *testing.T) {
	testCases := []struct {
		testCaseName   string
		testInput      ProjectConfig
		expectedErrors int
	}{
		{testCaseName: "Test case where the input is valid", testInput: ProjectConfig{
			projectName: "SampleProjectName",
			projectPath: "FilePath",
			projectRepo: "RepoRUL",
			projectType: false,
		}, expectedErrors: 0},
		{testCaseName: "Test case where projectName is missing", testInput: ProjectConfig{
			projectName: "",
			projectPath: "FilePath",
			projectRepo: "RepoRUL",
			projectType: false,
		}, expectedErrors: 1},
		{testCaseName: "Test case where projectPath is missing", testInput: ProjectConfig{
			projectName: "",
			projectPath: "FilePath",
			projectRepo: "RepoRUL",
			projectType: false,
		}, expectedErrors: 1},
		{testCaseName: "Test case where all settings are missing", testInput: ProjectConfig{
			projectName: "",
			projectPath: "",
			projectRepo: "",
			projectType: false,
		}, expectedErrors: 3},
	}

	for _, testCase := range testCases {
		errors := validateConf(testCase.testInput)
		if len(errors) != testCase.expectedErrors {
			t.Errorf("Test case:%s -- expected errors: #%d - got #%d errors",
				testCase.testCaseName, testCase.expectedErrors, len(errors))
		}

	}
}

func Test_generateScaffold(t *testing.T) {
	expectedStr := "🚀🚀 Generating scaffold for project hello-world in ./projects\n"
	buffer := bytes.Buffer{}
	generateScaffold(&buffer, ProjectConfig{projectName: "hello-world", projectPath: "./projects",
		projectRepo: "some url"})
	got := buffer.String()
	if !strings.EqualFold(got, expectedStr) {
		t.Errorf(got, expectedStr)
	}
}
