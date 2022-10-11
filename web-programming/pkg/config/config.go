package config

import "html/template"

// AppConfig holds the app configuration settings
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
