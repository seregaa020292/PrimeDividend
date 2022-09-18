package tpl

import (
	"github.com/flosch/pongo2/v6"
)

type Templater interface {
	Render(filepath string, vars map[string]any) ([]byte, error)
}

type template struct {
	useCache   bool
	globalVars map[string]any
	templates  *pongo2.TemplateSet
}

func NewTemplate(baseDir string, useCache bool, globalVars map[string]any) Templater {
	loader := pongo2.MustNewLocalFileSystemLoader(baseDir)
	templates := pongo2.NewSet("html", loader)

	return template{
		useCache:   useCache,
		globalVars: globalVars,
		templates:  templates,
	}
}

func (t template) Render(filepath string, vars map[string]any) ([]byte, error) {
	var (
		template *pongo2.Template
		err      error
	)

	if t.useCache {
		template, err = t.templates.FromCache(filepath)
	} else {
		template, err = t.templates.FromFile(filepath)
	}
	if err != nil {
		return nil, err
	}

	outputVars := t.copyGlobalVars()
	for k, v := range vars {
		outputVars[k] = v
	}

	return template.ExecuteBytes(outputVars)
}

func (t template) copyGlobalVars() map[string]any {
	copyVars := make(map[string]any)

	for k, v := range t.globalVars {
		copyVars[k] = v
	}

	return copyVars
}

func (t template) RegisterFilter(name string, fn pongo2.FilterFunction) error {
	if pongo2.FilterExists(name) {
		return pongo2.ReplaceFilter(name, fn)
	}

	return pongo2.RegisterFilter(name, fn)
}
