package tpl

import (
	"log"

	"github.com/aymerick/douceur/inliner"
	"github.com/flosch/pongo2/v6"
)

type Templater interface {
	Render(filepath string, vars map[string]any) (string, error)
	RenderInline(filepath string, vars map[string]any) (string, error)
}

type template struct {
	useCache   bool
	globalVars map[string]any
	templates  *pongo2.TemplateSet
}

func NewTemplate(baseDir string, useCache bool, globalVars map[string]any) Templater {
	loader, err := pongo2.NewLocalFileSystemLoader(baseDir)
	if err != nil {
		log.Fatalln(err)
	}

	templates := pongo2.NewSet("html", loader)

	return template{
		useCache:   useCache,
		globalVars: globalVars,
		templates:  templates,
	}
}

func (t template) Render(filepath string, vars map[string]any) (string, error) {
	var (
		tpl *pongo2.Template
		err error
	)

	if t.useCache {
		tpl, err = t.templates.FromCache(filepath)
	} else {
		tpl, err = t.templates.FromFile(filepath)
	}
	if err != nil {
		return "", err
	}

	outputVars := t.copyGlobalVars()
	for k, v := range vars {
		outputVars[k] = v
	}

	bytes, err := tpl.ExecuteBytes(outputVars)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (t template) RenderInline(filepath string, vars map[string]any) (string, error) {
	render, err := t.Render(filepath, vars)
	if err != nil {
		return "", err
	}

	return inliner.Inline(render)
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
