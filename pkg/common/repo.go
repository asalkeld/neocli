package repo

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const stack_template_repo = "https://github.com/nitrictech/stack-templates"

// Serialized Representations (YAML)
type TemplateDescriptor struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type RepositoryDescriptor struct {
	Templates []TemplateDescriptor `yaml:"templates"`
}

type Template struct {
	Repo *Repository
	Name string
	Path string
}

type Repository struct {
	Name      string
	Path      string
	Templates []Template
}

// HasTemplate - returns true if the repository contains a template with the name provided
func (r *Repository) HasTemplate(name string) bool {
	for _, t := range r.Templates {
		if t.Name == name {
			return true
		}
	}
	return false
}

func FromFile(file string) (*Repository, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	repoDes := &RepositoryDescriptor{}
	err = yaml.Unmarshal(yamlFile, repoDes)
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		Name:      filepath.Base(file),
		Path:      filepath.Dir(file),
		Templates: make([]Template, len(repoDes.Templates)),
	}

	for i, t := range repoDes.Templates {
		repo.Templates[i] = Template{
			Name: t.Name,
			Path: t.Path,
			Repo: repo,
		}
	}

	return repo, nil
}
