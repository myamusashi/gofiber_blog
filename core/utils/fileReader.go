package utils

import (
	"html/template"
	"io"
	"os"
)

type PostData struct {
	Title                   string   `yaml:"Title"`
	Slug                    string   `yaml:"Slug"`
	Date                    string   `yaml:"Date"`
	Description             string   `yaml:"Description"`
	Tags                    []string `yaml:"tags"`
	MetaDescription         string   `yaml:"MetaDescription"`
	MetaPropertyTitle       string   `yaml:"MetaPropertyTitle"`
	MetaPropertyDescription string   `yaml:"MetaPropertyDescription"`
	MetaOgURL               string   `yaml:"MetaOgURL"`
	Author                  Author   `yaml:"author"`
	Content                 template.HTML
}

type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type SlugRender interface {
	Read(slug string) (string, error)
}

// FileReader mengimplematasikan SlugRender interface
// Ini akan membaca/menyimpan isi data file yang dibaca
type FileReader struct{}

func (fRead FileReader) Read(slug string) (string, error) {
	fileRead, err := os.Open("markdown/" + slug + ".md")
	if err != nil {
		return "", err
	}
	defer fileRead.Close()
	b, err := io.ReadAll(fileRead)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
