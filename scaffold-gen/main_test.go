package main

import (
	"log"
	"os"
	"reflect"
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
